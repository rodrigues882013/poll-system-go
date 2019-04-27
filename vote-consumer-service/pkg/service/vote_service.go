package service

import (
	"context"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/commons"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/domain/models"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func NewVoteService(db *mongo.Client) *VoteService {

	return &VoteService{
		voteRepository:   repository.NewVoteRepository(db),
	}
}

type VoteService struct {
	voteRepository     repository.VoteRepository
	httpClient         *http.Client
}

func (v *VoteService) Create(context context.Context, vote models.Vote) (models.Vote, error) {
	vote.CreatedAt = time.Now()
	payload, err := v.voteRepository.Create(context, vote)
	commons.FailOnError(err, "Cannot be possible register vote")
	return payload, nil
}

