package infrastructure

import (
	"context"
	"github.com/felipe_rodrigues/poll-api/config"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type DB struct {
	Client *mongo.Client
}

var dbConn = &DB{}

func StartDB(conf config.Configuration) (*DB, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DataSourceURI))
	commons.FailOnError(err, "Failed to connect with MongoDB")

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	commons.FailOnError(err, "Failed to connect with MongoDB")

	dbConn.Client = client

	return dbConn, nil
}
