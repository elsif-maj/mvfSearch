package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/elsif-maj/umbraSearch/models"
)

func ConnectDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONN_STR"))
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	fmt.Println("Database connection established")
	return conn, nil
}

// Do I need to return an error here?
func GetAllSnippets(db *pgx.Conn) ([]models.Snippet, error) {
	queryResult, err := db.Query(context.Background(), "select * from \"Snippets\"")
	if err != nil {
		return nil, fmt.Errorf("Error with DB Query: %v", err)
	}

	if queryResult == nil {
		return nil, fmt.Errorf("queryResult is nil")
	}

	var allSnippets = []models.Snippet{}

	for queryResult.Next() {
		var s models.Snippet
		err = queryResult.Scan(&s.Id, &s.Title, &s.Language, &s.Code, &s.UserId, &s.Created_at, &s.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Problem with queryResult.Scan() step: %v", err)
		}
		allSnippets = append(allSnippets, s)
	}

	return allSnippets, nil
}
