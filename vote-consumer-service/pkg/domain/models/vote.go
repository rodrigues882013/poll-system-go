package models

import "time"

type Vote struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Nominate    Nominate `json:"nominate"`
	Poll        Poll     `json:"poll"`
	CreatedAt   time.Time `json:"createdAt"`
}