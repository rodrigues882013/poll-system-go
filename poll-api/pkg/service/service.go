package service

import (
	"context"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
)

type PollService interface {
	FindById(ctx context.Context, id int64) (*models.Poll, error)
	Create(ctx context.Context, poll models.Poll) (*models.Poll, error)
	Count(ctx context.Context, id int64) (*models.PollResult, error)
	CountByNominate(ctx context.Context, id int64) (*models.PollResultByNominates, error)
	CountByHour(ctx context.Context, id int64) ([]models.PollResultByHour, error)
	IsValidPoll(ctx context.Context, poll models.Poll) bool
	IsPollClosed(ctx context.Context, poll models.Poll) bool

}

type VoteService interface {
	Vote(vote models.Vote, queueName string) <- chan bool
	CanVote(ctx context.Context, params map[string]string, vote models.Vote) (bool, models.GeneralResponse, int)
}
