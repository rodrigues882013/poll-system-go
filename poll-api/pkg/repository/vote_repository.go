package repository

import (
	"context"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

	if err != nil {
		log.Fatal(err)
	}

	return vote, err
}




