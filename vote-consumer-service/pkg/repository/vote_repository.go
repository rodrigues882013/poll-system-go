package repository

import (
	"context"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/commons"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewVoteRepository(Conn *mongo.Client) VoteRepository {
	return &voteRepository {
		Conn: Conn,
	}
}

type voteRepository struct {
	Conn *mongo.Client
}

func (v *voteRepository) Create(ctx context.Context, vote models.Vote) (models.Vote, error) {
	collection := v.Conn.Database("pollapi").Collection("votes")
	_, err := collection.InsertOne(ctx, vote)
	commons.FailOnError(err, "Cannot was possible crate votes")
	return vote, err
}




