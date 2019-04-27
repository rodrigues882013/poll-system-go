package repository

import (
	"encoding/json"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

func NewPollCacheRepository(cacheClient *redis.Client) PollCacheRepository {
	return &pollCacheRepository{
		CacheClient: cacheClient,
	}
}

type pollCacheRepository struct {
	CacheClient  *redis.Client
}

func (c *pollCacheRepository) Get(k int64) (*models.Poll, error){
	key := strconv.FormatInt(int64(k), 16)
	val, err := c.CacheClient.Get(key).Result()

	if err != nil {
		return nil, err
	}
	poll := &models.Poll{}
	err = json.Unmarshal([]byte(val), &poll)

	return poll, err
}

func (c *pollCacheRepository) Set(k int64, data *models.Poll){
	b, err := json.Marshal(data)
	commons.FailOnError(err, "Cannot be possible marshaling the data")
	key := strconv.FormatInt(int64(k), 16)
	err1 := c.CacheClient.Set(key, b, time.Duration(3600000000000 * data.Duration)).Err()
	commons.FailOnError(err1, "Cannot be possible cached the data")
	log.Println("Cached with success")
}
