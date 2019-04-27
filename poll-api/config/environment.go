package config

type Configuration struct {
	Context               string        `env:"CONTEXT" envDefault:"/poll-api"`
	Port                  int           `env:"PORT" envDefault:"3000"`
	QueueHost             string        `env:"QUEUE_HOST" envDefault:"amqp://guest:guest@localhost:5672/"`
	QueueName             string        `env:"QUEUE_NAME" envDefault:"votequeue"`
	DataSourceURI         string        `env:"DATA_SOURCE_URI" envDefault:"mongodb://localhost:27017"`
	RedisHost             string        `env:"REDIS_HOST" envDefault:"localhost:6379"`
}