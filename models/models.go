package models

import "time"

type Snippet struct {
	Id        int
	Title     string
	Language  string
	Code      string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
