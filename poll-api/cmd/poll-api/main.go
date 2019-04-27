package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/felipe_rodrigues/poll-api/config"
	infrastructure "github.com/felipe_rodrigues/poll-api/pkg/infrastructure"
	"github.com/felipe_rodrigues/poll-api/pkg/handler"
	"log"
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

	cacheClient := infrastructure.CreateCacheClient(cfg)

	handler.Init(cfg, client, conn, cacheClient)

}
