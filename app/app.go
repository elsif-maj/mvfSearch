package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/elsif-maj/umbraSearch/indexer"
	"github.com/elsif-maj/umbraSearch/myEnv"
	"github.com/jackc/pgx/v5"
)

type Server struct {
	DevTestIndex map[string][]int // Development data structure for testing routes, tokenization, n-gramification, etc.
	DBConn       *pgx.Conn
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func Setup() (*Server, error) {
	// Set environment variables if needed
	myEnv.SetEnv()

	// Connect to DB
	dbConn, err := db.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	app := &Server{
		DBConn: dbConn,
	}

	return app, nil
}

func MakeAPIFunc(fn apiFunc) http.HandlerFunc {
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

func (app *Server) HandleSnippets(w http.ResponseWriter, r *http.Request) error {
	snippets, err := db.GetAllSnippets(app.DBConn)
	if err != nil {
		return fmt.Errorf("Problem with db.GetAllSnippets: %v", err)
	}

	return writeJSON(w, http.StatusOK, snippets)
}

// Test:
// curl -X POST -H "Content-Type: application/json" -d '{"id":16}' http://localhost:3000/api/snippets/new
func (app *Server) HandleNewSnippet(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
	}

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("Problem with decoding request body: %v", err)
	}

	id, ok := data["id"].(float64)
	if !ok {
		return fmt.Errorf("Invalid or missing 'id' in request body JSON object")
	}

	snippetId := int(id)

	indexer.ProcessInput(app.DBConn, snippetId)

	return writeJSON(w, http.StatusOK, map[string]string{"Success": fmt.Sprintf("The creation of snippet with DB primary key id: %d has been registered.", snippetId)})
}
