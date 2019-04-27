package http

import (
	"github.com/felipe_rodrigues/poll-api/config"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	"github.com/felipe_rodrigues/poll-api/pkg/infrastructure"
	"github.com/felipe_rodrigues/poll-api/pkg/service"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	voteService service.VoteService
	queueName string
)

func InitVoteHandler(conf config.Configuration,
	router *mux.Router,
	database *infrastructure.DB,
	broker *infrastructure.Broker,
	client *redis.Client){

	voteService = service.NewVoteService(database.Client, client, broker.Conn)

	queueName = conf.QueueName
	router.HandleFunc(conf.Context + "/polls/{pollId}/vote", RegisterVote).Methods("POST")
}

func RegisterVote(w http.ResponseWriter, request *http.Request){
	vote := models.Vote{}
	commons.BindJSON(w, request, &vote)
	params := mux.Vars(request)

	canVote, message, status := voteService.CanVote(request.Context(), params, vote)

	if canVote {
		voteService.Vote(vote, queueName)
	}

	commons.Render(w, status, &message)
}

