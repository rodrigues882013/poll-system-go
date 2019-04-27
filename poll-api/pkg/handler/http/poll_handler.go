package http

import (
	"github.com/felipe_rodrigues/poll-api/config"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/felipe_rodrigues/poll-api/pkg/domain/models"
	database "github.com/felipe_rodrigues/poll-api/pkg/infrastructure"
	"github.com/felipe_rodrigues/poll-api/pkg/service"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	serv     service.PollService
	XBBBRoot string = "X-BBB-ROOT"
)

func InitPollHandler(conf config.Configuration, router *mux.Router, database *database.DB, client *redis.Client){
	serv = service.NewPollService(database.Client, client)

	router.HandleFunc(conf.Context + "/polls", SavePoll).Methods("POST")
	router.HandleFunc(conf.Context + "/polls/{pollId}", GetPollById).Methods("GET")
	router.HandleFunc(conf.Context + "/polls/{pollId}/results", GetResultByPollIdAndNominate).
		Queries("byNominate", "{byNominate}").
		Methods("GET")

	router.HandleFunc(conf.Context + "/polls/{pollId}/results", GetResultByPollIdHour).
		Queries("byHour", "{byHour}").
		Methods("GET")

	router.HandleFunc(conf.Context + "/polls/{pollId}/results", GetResultByPollId)


}

func GetResultByPollId(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pollId, err := strconv.Atoi(params["pollId"])

	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	_, err = serv.FindById(request.Context(), int64(pollId))
	if err != nil {
		commons.Render(w, http.StatusNotFound, &models.GeneralResponse{Message: "Poll not found"})
		return
	}

	re, err := serv.Count(request.Context(), int64(pollId))
	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	commons.Render(w, http.StatusOK, re)

}

func GetResultByPollIdHour(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pollId, err := strconv.Atoi(params["pollId"])
	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	_, err = serv.FindById(request.Context(), int64(pollId))
	if err != nil {
		commons.Render(w, http.StatusNotFound, &models.GeneralResponse{Message: "Poll not found"})
		return
	}

	re, err := serv.CountByHour(request.Context(), int64(pollId))
	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	commons.Render(w, http.StatusOK, re)

}

func GetResultByPollIdAndNominate(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pollId, err := strconv.Atoi(params["pollId"])
	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	_, err = serv.FindById(request.Context(), int64(pollId))
	if err != nil {
		commons.Render(w, http.StatusNotFound, &models.GeneralResponse{Message: "Poll not found"})
		return
	}

	re, err := serv.CountByNominate(request.Context(), int64(pollId))

	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	commons.Render(w, http.StatusOK, re)
}

func GetPollById(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	pollId, err := strconv.Atoi(params["pollId"])

	if err != nil {
		log.Println(err)
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "Bad request"})
		return
	}

	poll, err := serv.FindById(request.Context(), int64(pollId))
	if err != nil {
		commons.Render(w, http.StatusNotFound, &models.GeneralResponse{Message: "Poll not found"})
		return
	}

	if serv.IsPollClosed(request.Context(), *poll) {

		//Após o fechamento de uma votação só permito acesso a quem passa o header especial
		if request.Header.Get(XBBBRoot) != "ROOT" {
			commons.Render(w, http.StatusForbidden, &models.GeneralResponse{Message: "The poll is closed, you cannot access."})
			return
		}
	}

	commons.Render(w, http.StatusOK, &poll)
}

func SavePoll(w http.ResponseWriter, request *http.Request) {
	var poll models.Poll

	if !commons.BindJSON(w, request, &poll) || serv.IsValidPoll(request.Context(), poll) {
		commons.Render(w, http.StatusBadRequest, &models.GeneralResponse{Message: "You cannot create poll."})
		return
	}

	p, err := serv.Create(request.Context(), poll)
	if err != nil {
		status, message := http.StatusNotAcceptable, &models.GeneralResponse{Message: "Poll cannot be created"}

		if err.Error() == "conflict" {
			status, message = http.StatusConflict, &models.GeneralResponse{Message: "Poll with id already exist"}
		}

		commons.Render(w, status, message)
		return
	}

	commons.Render(w, http.StatusCreated, p)
}


