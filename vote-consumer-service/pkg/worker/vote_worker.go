package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/felipe_rodrigues/vote-consumer-service/config"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/commons"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/domain/models"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/infrastructure"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/service"
	"log"
	"time"
)


var voteService *service.VoteService
var conn *infrastructure.Broker
var queueName string

func InitVoteWorker(conf config.Configuration, database *infrastructure.DB,
	broker *infrastructure.Broker){
	voteService = service.NewVoteService(database.Client)
	conn = broker
	queueName = conf.QueueName
	consumeVote()
}

func consumeVote() {
	if conn.Conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	// One channel for each go routine, only reuse connection (one for each process)
	ch, err := conn.Conn.Channel()
	commons.FailOnError(err, "Connection cannot be established")
	defer ch.Close()

	// Publishes a message onto the queue.
	msgs, err := ch.Consume(
		queueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	commons.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var vote models.Vote
			err := json.Unmarshal(d.Body, &vote)
			commons.FailOnError(err, "Message cannot be unmarshal")

			_, err = voteService.Create(context.Background(), vote)
			commons.FailOnError(err, "Vote wasn't computed.")

			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)

			if err == nil {
				log.Printf("Vote was computed with success")
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


}
