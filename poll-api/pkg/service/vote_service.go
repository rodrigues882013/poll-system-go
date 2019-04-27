package service

import (
	"context"
	"encoding/json"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/felipe_rodrigues/poll-api/pkg/repository"
	"github.com/felipe_rodrigues/poll-api/pkg/route"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func NewVoteService(db *mongo.Client, redis *redis.Client, conn *amqp.Connection) VoteService {
	return &voteService {
		Repository: repository.NewVoteRepository(db),
		Route: route.NewVoteRoute(conn),
		PollService: NewPollService(db, redis),
	}
}

type voteService struct {
	Repository   repository.VoteRepository
	Route        route.VoteRoute
	PollService  PollService
}

func (v *voteService) Vote(vote models.Vote, queueName string) <- chan bool {
	c := make(chan bool)

	go func() {
		b, _ := json.Marshal(vote)
		isOK, err := v.Route.PublishVote([]byte(b), queueName)
		commons.FailOnError(err, "Item cannot be sent to queue")

		if isOK {
			log.Println("Item was sent with success")
			c <- true
		} else {
			c <- false
		}
	}()

	return c

}

func (v *voteService) CanVote(ctx context.Context, params map[string]string, vote models.Vote) (bool, models.GeneralResponse, int) {

	pollId, err := strconv.Atoi(params["pollId"])

	if err != nil {
		return false, models.GeneralResponse{Message: "This pollId was not given."}, http.StatusBadRequest
	}

	if int64(pollId) != int64(vote.Poll.Id) {
		return false, models.GeneralResponse{Message: "This poll doesn't exist."}, http.StatusPreconditionFailed
	}

	// Preciso validar se o voto é valido, esse dado é cachecado
	poll, err := v.PollService.FindById(ctx, int64(vote.Poll.Id))

	if err != nil {
		return false, models.GeneralResponse{Message: "This poll doesn't exist."}, http.StatusNotFound
	}

	if poll.Nominates == nil || len(poll.Nominates) == 0 {
		return false, models.GeneralResponse{Message: "You do not chosen any nominate."}, http.StatusPreconditionFailed
	}

	if !commons.InSlice(vote.Nominate, poll.Nominates) {
		return false, models.GeneralResponse{Message: "The nominate there isn't belong this poll."}, http.StatusPreconditionFailed
	}

	if commons.IsOutOfPeriod(poll.StartAt, int64(poll.Duration)) {
		return false, models.GeneralResponse{Message: "The poll is closed."}, http.StatusForbidden
	}

	return true, models.GeneralResponse{Message: "Your vote was registered with successful."}, http.StatusAccepted
}
