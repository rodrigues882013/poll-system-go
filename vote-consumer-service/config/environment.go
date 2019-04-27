package config

type Configuration struct {
	Context               string        `env:"CONTEXT" envDefault:"/vote-consumer-service"`
	Port                  int           `env:"PORT" envDefault:"3001"`
	QueueHost             string        `env:"QUEUE_HOST" envDefault:"amqp://guest:guest@localhost:5672/"`
	QueueName             string        `env:"QUEUE_NAME" envDefault:"votequeue"`
	DataSourceURI         string        `env:"DATA_SOURCE_URI" envDefault:"mongodb://localhost:27017"`
}