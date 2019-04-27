package infrastructure

import (
	"github.com/felipe_rodrigues/poll-api/config"
	"github.com/felipe_rodrigues/poll-api/pkg/commons"
	"github.com/streadway/amqp"
)

type Broker struct {
	Conn *amqp.Connection
}

var brokerConn = &Broker{}


func StartBroker(conf config.Configuration) (*Broker, error) {
	conn, err := amqp.Dial(conf.QueueHost)
	commons.FailOnError(err, "Failed to open a channel")

	if err != nil {
		panic("Application cannot continue without broker connection")
	}
	brokerConn.Conn = conn

	return brokerConn, nil
}
