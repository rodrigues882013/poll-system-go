package route

type VoteRoute interface {
	PublishVote(body []byte, queueName string) (bool, error)
}
