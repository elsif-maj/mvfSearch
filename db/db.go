package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

/*
Things to eventually change about the inherited schema:
  - Relation names/titles that feature uppercase letters are against best practice and necessiate using quotation marks to specify the table in all queries.
  - There is a mix of camelCase and snake_case in the column names below
*/
type Snippet struct {
	id         int
	title      string
	language   string
	code       string
	userId     string
	created_at time.Time
	updatedAt  time.Time
}

func ConnectDB() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), os.Getenv("DB_CONN_STR"))
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	return db, nil
}

// queryResult, err := db.Query(context.Background(), "select * from \"Snippets\"")
// if err != nil {
// 	log.Fatal("Error with DB Query: ", err)
// }

// var allSnippets = []Snippet{}

// for queryResult.Next() {
// 	var s Snippet
// 	err = queryResult.Scan(&s.id, &s.title, &s.language, &s.code, &s.userId, &s.created_at, &s.updatedAt)
// 	if err != nil {
// 		fmt.Println("Problem with queryResult.Scan() step: ", err)
// 	}
// 	allSnippets = append(allSnippets, s)
// }

// fmt.Println(allSnippets[len(allSnippets)-1])
