package service

import (
	"context"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/felipe_rodrigues/poll-api/pkg/repository"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func NewPollService(client *mongo.Client, redis *redis.Client) PollService {
	return &pollService{
		PollRepository: repository.NewPollRepository(client),
		PollCacheRepository: repository.NewPollCacheRepository(redis),
	}
}

type pollService struct {
	PollRepository repository.PollRepository
	PollCacheRepository repository.PollCacheRepository
}

func (p *pollService) IsPollClosed(ctx context.Context, poll models.Poll) bool {
	return commons.IsOutOfPeriod(poll.StartAt, int64(poll.Duration))
}

func findById(repo repository.PollRepository, cache repository.PollCacheRepository, id int64) (*models.Poll, error) {
	poll, err := cache.Get(id)

	if err != nil {
		err = nil

		poll, err := repo.FindById(context.Background(), id)

		if err == nil {
			cache.Set(id, poll)
			return &models.Poll{Id: poll.Id, Nominates:poll.Nominates, Year:poll.Year, Duration:poll.Duration}, err

		} else {
			return nil, commons.NotFoundError{Message: "Poll not found"}
		}
	}

	return poll, err
}

func (p *pollService) IsValidPoll(ctx context.Context, poll models.Poll) bool{
	return poll.Nominates == nil || len(poll.Nominates) == 0 || len(poll.Nominates) == 1
}

func (p *pollService) FindById(ctx context.Context, id int64) (*models.Poll, error) {
	return findById(p.PollRepository, p.PollCacheRepository, id)
}

func (p *pollService) Count(ctx context.Context, id int64) (*models.PollResult, error) {
	return p.PollRepository.Count(ctx, id)
}

func (p *pollService) CountByHour(ctx context.Context, id int64) ([]models.PollResultByHour, error) {
	return p.PollRepository.CountByHours(ctx, id)
}

func (p *pollService) CountByNominate(ctx context.Context, id int64) (*models.PollResultByNominates, error) {
	return p.PollRepository.CountByNominates(ctx, id)
}

func (p *pollService) Create(ctx context.Context, poll models.Poll) (*models.Poll, error)  {

	pollAlreadyExist, _ := findById(p.PollRepository, p.PollCacheRepository, int64(poll.Id))

	if pollAlreadyExist != nil {
		return nil, commons.ConflictedError{Message:"conflict"}
	}

	if poll.StartAt == *new(time.Time){
		poll.StartAt = time.Now()
	}

	return p.PollRepository.Create(ctx, poll)
}
