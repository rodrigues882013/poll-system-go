package models

type Poll struct {
	Id               int        `json:"id"`
	Year             int        `json:"year"`
	Duration         int        `json:"duration"`
	Nominates        []Nominate `json:"nominates"`
}

type PollResultEntry struct {
	Nominates Nominate `json:"nominate"`
	Votes     int      `json:"votes"`
}

type PollResult struct {
	PollId        int                 `json:"pollId"`
	Result        int64               `json:"results"`
}

type PollResultByNominates struct {
	PollId        int                 `json:"pollId"`
	Result        []PollResultEntry   `json:"results"`
}

type PollResultByHour struct {
	PollId        int                 `json:"pollId"`
	Result        []PollResultEntry   `json:"results"`
	Hour          string              `json:"hour"`
}
