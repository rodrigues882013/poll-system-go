package infrastructure

import (
	"github.com/felipe_rodrigues/poll-api/config"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/go-redis/redis"
)

func CreateCacheClient(configuration config.Configuration) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     configuration.RedisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	commons.FailOnError(err, "Error on connect with redis")

	return client
}