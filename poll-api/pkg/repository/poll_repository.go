package repository

import (
	"context"
	"fmt"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
)


func NewPollRepository(client *mongo.Client) PollRepository {
	return &pollRepository {
		Client: client,
	}
}

type pollRepository struct {
	Client *mongo.Client
}

func toInt(data interface{}) int64{
	val := fmt.Sprintf("%v", data)
	totalVotes, err := strconv.ParseInt(val, 10, 64)

	if err != nil {
		log.Println(err)
	}

	return totalVotes
}

func (p *pollRepository) Count(ctx context.Context, id int64) (*models.PollResult, error) {
	collection := p.Client.Database("pollapi").Collection("votes")
	pipeline := mongo.Pipeline{
		{{
			Key: "$match", Value: bson.D{{Key: "poll.id", Value: id}},
		}},
		{{
			Key: "$count", Value: "totalVotes",
		}},
	}
	cur, err := collection.Aggregate(ctx, pipeline)
	var re *models.PollResult

	if err != nil {
		log.Println(err)

	} else {
		var itemMap map[string]interface{}

		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)

			if err != nil {
				log.Println(err)
			}

			b, _ := bson.Marshal(result)
			err = bson.Unmarshal(b, &itemMap)
			re = &models.PollResult{PollId: id, Total: toInt(itemMap["totalVotes"])}

		}
	}

	return re, err
}

func (p *pollRepository) CountByNominates(ctx context.Context, id int64) (*models.PollResultByNominates, error) {
	collection := p.Client.Database("pollapi").Collection("votes")
	var rest *models.PollResultByNominates

	pipeline := mongo.Pipeline{
		{{
			Key: "$match", Value: bson.D{{Key: "poll.id", Value: id}},
		}},
		{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "nominate", Value: "$nominate.id"}}},
				{Key:"count", Value: bson.D{{Key: "$sum", Value: 1}}}},
		}},
	}

	cur, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		log.Println(err)

	} else {
		var items []interface{}

		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)

			if err != nil {
				log.Println(err)
			}
			items = append(items, result)
		}

		var itemMap map[string]interface{}
		var innerMap map[string]interface{}
		var entries []models.PollResultEntry

		for i := 0; i < len(items); i++ {
			b, _ := bson.Marshal(items[i])
			err = bson.Unmarshal(b, &itemMap)
			val := itemMap["_id"]
			b, _ = bson.Marshal(val)
			err = bson.Unmarshal(b, &innerMap)

			entry := models.PollResultEntry{Nominates: models.Nominate{Id: toInt(innerMap["nominate"])},
				Votes: toInt(itemMap["count"])}

			entries = append(entries, entry)
		}

		rest = &models.PollResultByNominates{PollId: int(id), Result: entries}

	}

	return rest, nil
}

func (p *pollRepository) CountByHours(ctx context.Context, id int64) ([]models.PollResultByHour, error) {
	collection := p.Client.Database("pollapi").Collection("votes")
	var entries []models.PollResultByHour

	pipeline := mongo.Pipeline{
		{{
			Key: "$match", Value: bson.D{{Key: "poll.id", Value: id}},
		}},
		{{
			Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "month", Value: bson.D{{ Key: "$month", Value: "$createdat"}}},
					{Key:"day", Value: bson.D{{ Key: "$dayOfMonth", Value: "$createdat"}}},
					{Key:"year", Value: bson.D{{ Key: "$year", Value: "$createdat"}}},
					{Key:"hour", Value: bson.D{{ Key: "$hour", Value: "$createdat"}}}}},
				{Key:"count", Value: bson.D{{Key: "$sum", Value: 1}}}},
		}},
	}

	cur, err := collection.Aggregate(ctx, pipeline)

	if err != nil {
		log.Println(err)

	} else {
		var items []interface{}

		for cur.Next(ctx) {
			var result bson.M
			err := cur.Decode(&result)

			if err != nil {
				log.Println(err)
			}
			items = append(items, result)
		}

		var itemMap map[string]interface{}
		var innerMap map[string]interface{}

		for i := 0; i < len(items); i++ {
			b, _ := bson.Marshal(items[i])
			err = bson.Unmarshal(b, &itemMap)

			val := itemMap["_id"]
			b, _ = bson.Marshal(val)
			err = bson.Unmarshal(b, &innerMap)
			entry := models.PollResultByHour{PollId: int64(id), Hour: toInt(innerMap["hour"]), Votes: toInt(itemMap["count"])}
			entries = append(entries, entry)
		}

	}

	return entries, nil
}

func (p *pollRepository) FindById(ctx context.Context, id int64) (*models.Poll, error) {
	collection := p.Client.Database("pollapi").Collection("polls")
	filter := bson.M{"id": id}

	var poll models.Poll

	err := collection.FindOne(ctx, filter).Decode(&poll)

	if err != nil {
		log.Println(err)
	}

	return &poll, err

}

func (p *pollRepository) Create(ctx context.Context, poll models.Poll) (*models.Poll, error) {
	collection := p.Client.Database("pollapi").Collection("polls")

	_, err := collection.InsertOne(ctx, poll)

	if err != nil {
		log.Println(err)
	}

	return &poll, err
}




