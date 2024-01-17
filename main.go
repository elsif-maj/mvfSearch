package main

import (
	// std
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// dependencies
	"github.com/jackc/pgx/v5"
	// local imports
	"github.com/elsif-maj/umbraSearch/myEnv"
	"github.com/elsif-maj/umbraSearch/tokenizer"
)

type User struct {
	Username string
}

func main() {
	// Set environment variables if needed
	myEnv.SetEnv()

	// DB
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONN_STR"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queryResult, err := conn.Query(context.Background(), "select * from public.\"Snippets\"")
	if err != nil {
		log.Fatal("Error with DB Query: ", err)
	}

	var col2, col3, col4 string
	var col1, col5 int
	var col6, col7 time.Time

	for queryResult.Next() {
		err = queryResult.Scan(&col1, &col2, &col3, &col4, &col5, &col6, &col7)
		if err != nil {
			fmt.Println("Problem with queryResult.Scan() step: ", err)
			break
		}
		fmt.Println("Result from PG Database Query: ", col2) //col1, col2, col3, col4, col5, col6, col7)
	}

	// Test & Temp
	fmt.Println(tokenizer.Tokenize("Tokenize this bro! testing what's beyond the punctuation..."))

	http.HandleFunc("/", makeAPIFunc(handleHome))         // Placeholder/Test route
	http.HandleFunc("/api/user", makeAPIFunc(handleUser)) // Placeholder/Test route

	// Main Server Loop
	http.ListenAndServe(":3000", nil)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) error {
	user := User{
		Username: "Hassan",
	}
	return writeJSON(w, http.StatusOK, map[string]string{"user": user.Username})
}

func handleUser(w http.ResponseWriter, r *http.Request) error {
	// return fmt.Errorf("failed to do something")
	return writeJSON(w, http.StatusOK, map[string]string{"message": "hello some user"})
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
