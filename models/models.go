package models

import "time"

/*
Things to eventually change about the inherited schema:
  - Relation names/titles that feature uppercase letters are against best practice and necessiate using quotation marks to specify the table in all queries.
  - There is a mix of camelCase and snake_case in the column names below
*/
type Snippet struct {
	Id         int
	Title      string
	Language   string
	Code       string
	UserId     string
	Created_at time.Time
	UpdatedAt  time.Time
}
