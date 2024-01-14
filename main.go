package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elsif-maj/umbraSearch/tokenizer"
)

type User struct {
	Username string
}

func main() {
	fmt.Println(tokenizer.Tokenize("Tokenize this bro! testing what's beyond ok?"))

	http.HandleFunc("/", makeAPIFunc(handleHome))
	http.HandleFunc("/api/user", makeAPIFunc(handleUser))
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
