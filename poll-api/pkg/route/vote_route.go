package route

import (
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/streadway/amqp"
)

func NewVoteRoute(conn *amqp.Connection) VoteRoute {
	return &voteRoute {
		Conn: conn,
	}
}

type voteRoute struct {
	Conn *amqp.Connection
}

func (v *voteRoute) PublishVote(body []byte, queueName string) (bool, error) {
	if v.Conn == nil {
		panic("Tried to send message before connection was initialized. Don't do that.")
	}

	// One channel for each go routine, only reuse connection (one for each process)
	ch, err := v.Conn.Channel()
	commons.FailOnError(err, "Connection cannot be established")
	defer ch.Close()


	queue, err := ch.QueueDeclare(
		queueName, // our queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil, // arguments
	)

	commons.FailOnError(err, "Item cannot be sent to queue")


	// Publishes a message onto the queue.
	err = ch.Publish(
		"", // use the default exchange
		queue.Name, // routing key, e.g. our queue name
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body, // Our JSON body as []byte
		})

	return err == nil, err

}



