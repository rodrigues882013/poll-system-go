package route

type VoteRoute interface {
	ConsumeVote(queueName string)
}
