package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/felipe_rodrigues/vote-consumer-service/config"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/infrastructure"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/worker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)


func main(){

	// Load config
	cfg := config.Configuration{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	//Init
	client, err := infrastructure.StartDB(cfg)

	if err != nil {
		log.Fatal(err)
	}

	conn, err := infrastructure.StartBroker(cfg)

	worker.Init(cfg, client, conn)

	// Start
	go startWebContext(cfg)

}

func HealthEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func startWebContext(cfg config.Configuration){
	router := mux.NewRouter()
	router.HandleFunc(cfg.Context + "/health", HealthEndpoint).Methods("GET")
	log.Println("LISTENING ON PORT ", cfg.Port)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(cfg.Port), router))
}
