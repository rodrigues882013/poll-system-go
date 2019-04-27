package handler

import (
	"github.com/felipe_rodrigues/poll-api/config"
	handler "github.com/felipe_rodrigues/poll-api/pkg/handler/http"
	"github.com/felipe_rodrigues/poll-api/pkg/infrastructure"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func Init(conf config.Configuration, db *infrastructure.DB, ch *infrastructure.Broker, client *redis.Client) {
	router := mux.NewRouter()
	router.HandleFunc(conf.Context + "/health", HealthEndpoint).Methods("GET")

	// Init routers
	handler.InitPollHandler(conf, router, db, client)
	handler.InitVoteHandler(conf, router, db, ch, client)

	log.Println("LISTENING ON PORT ", conf.Port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(conf.Port),
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))
}


