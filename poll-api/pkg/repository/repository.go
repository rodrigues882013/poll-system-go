package repository

import (
	"context"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
)

type PollRepository interface {
	FindById(ctx context.Context, id int64) (*models.Poll, error)
	Create(ctx context.Context, data models.Poll) (*models.Poll, error)
	Count(ctx context.Context, id int64) (*models.PollResult, error)
	CountByNominates(ctx context.Context, id int64) (*models.PollResultByNominates, error)
	CountByHours(ctx context.Context, id int64) ([]models.PollResultByHour, error)
}

type PollCacheRepository interface {
	Get(k int64)(*models.Poll, error)
	Set(k int64, poll *models.Poll)
}

type VoteRepository interface {
	Create(ctx context.Context, vote models.Vote) (models.Vote, error)
}


