package models

import "time"

type Poll struct {
	Id               int        `json:"id"`
	Year             int        `json:"year"`
	Duration         int        `json:"duration"`
	Nominates        []Nominate `json:"nominates"`
	StartAt          time.Time  `json:"startAt"`
}

type PollResultEntry struct {
	Nominates Nominate `json:"nominate"`
	Votes     int64      `json:"votes"`
}

type PollResult struct {
	PollId        int64    `json:"pollId"`
	Total         int64    `json:"votes"`
}

type PollResultByNominates struct {
	PollId        int                 `json:"pollId"`
	Result        []PollResultEntry   `json:"results"`
}

type PollResultByHour struct {
	PollId        int64                 `json:"pollId"`
	Votes         int64               `json:"votes"`
	Hour          int64              `json:"hour"`
}
