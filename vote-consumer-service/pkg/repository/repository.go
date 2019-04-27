package repository

import (
	"context"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/domain/models"
)

type VoteRepository interface {
	Create(ctx context.Context, vote models.Vote) (models.Vote, error)
}


