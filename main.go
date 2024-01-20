package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/elsif-maj/umbraSearch/myEnv"
)

type User struct {
	Username string
}

func main() {
	// Set environment variables if needed
	myEnv.SetEnv()

	// DB
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	// Test & Temp
	// fmt.Println(tokenizer.Tokenize("Test token. Also: Testing what's beyond the punctuation..."))

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
