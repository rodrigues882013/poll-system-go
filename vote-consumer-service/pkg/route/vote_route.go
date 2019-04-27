package route

import (
	"bytes"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/commons"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func NewVoteRoute(conn *amqp.Connection) VoteRoute {
	return &voteRoute {
		Conn: conn,
	}
}

type voteRoute struct {
	Conn *amqp.Connection
}

func (v *voteRoute) ConsumeVote(queueName string) {
	if v.Conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	// One channel for each go routine, only reuse connection (one for each process)
	ch, err := v.Conn.Channel()
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
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever


}



