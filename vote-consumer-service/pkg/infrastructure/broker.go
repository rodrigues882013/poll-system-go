package infrastructure

import (
	"github.com/felipe_rodrigues/vote-consumer-service/config"
	"github.com/felipe_rodrigues/vote-consumer-service/pkg/commons"
	"github.com/streadway/amqp"
)

type Broker struct {
	Conn *amqp.Connection
}

var brokerConn = &Broker{}


func StartBroker(conf config.Configuration) (*Broker, error) {
	conn, err := amqp.Dial(conf.QueueHost)
	commons.FailOnError(err, "Failed to open a channel")
	brokerConn.Conn = conn

	return brokerConn, nil
}
