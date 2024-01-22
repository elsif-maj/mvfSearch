package db

import (
	"context"
	"errors"
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

func GetSnippet(dbconn *pgx.Conn, id int) (models.Snippet, error) {
	var s models.Snippet

	queryResult := dbconn.QueryRow(
		context.Background(),
		"SELECT * FROM \"Snippets\" WHERE id=$1",
		id,
	)

	err := queryResult.Scan(&s.Id, &s.Title, &s.Language, &s.Code, &s.UserId, &s.Created_at, &s.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return s, fmt.Errorf("No snippet found with ID %d", id)
		}
		return s, fmt.Errorf("Error with DB Query: %v", err)
	}
	return s, nil
}

func GetAllSnippets(db *pgx.Conn) ([]models.Snippet, error) {
	queryResult, err := db.Query(context.Background(), "SELECT * FROM \"Snippets\"")
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
