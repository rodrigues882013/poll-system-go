package worker

import (
	"github.com/felipe_rodrigues/vote-consumer-service/config"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/infrastructure"
)



func Init(conf config.Configuration, db *infrastructure.DB, ch *infrastructure.Broker) {
	InitVoteWorker(conf, db, ch)
}


