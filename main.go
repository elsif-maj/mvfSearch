package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/elsif-maj/umbraSearch/myEnv"
	"github.com/jackc/pgx/v5"
)

type apiFunc func(http.ResponseWriter, *http.Request) error
type App struct {
	DevTestIndex map[string][]int // Development data structure for testing routes, tokenization, n-gramification, etc.
	DBConn       *pgx.Conn
}

func main() {
	// Set environment variables if needed
	myEnv.SetEnv()

	// DB
	dbConn, err := db.ConnectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer dbConn.Close(context.Background())

	app := &App{
		DBConn: dbConn,
	}

	// Routes
	http.HandleFunc("/api/snippets", makeAPIFunc(app.handleSnippets)) // Placeholder/Test route

	// Main Server Loop
	http.ListenAndServe(":3000", nil)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func (app *App) handleSnippets(w http.ResponseWriter, r *http.Request) error {
	snippets, err := db.GetAllSnippets(app.DBConn)
	if err != nil {
		return fmt.Errorf("Problem with db.GetAllSnippets: %v", err)
	}

	return writeJSON(w, http.StatusOK, snippets)
}
