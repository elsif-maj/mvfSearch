package indexer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/jackc/pgx/v5"
)

func ProcessInput(dbconn *pgx.Conn, id int) error {
	// Get snippet from database
	snippet, err := db.GetSnippet(dbconn, id)
	if err != nil {
		return fmt.Errorf("failed to get snippet from database: %w", err)
	}

	// Tokenize snippet -- come back to this and add the title
	i, err := tokenize(snippet.Code)
	if err != nil {
		return fmt.Errorf("failed to tokenize snippet id: %d", id)
	}

	fmt.Println(i)
	/*
		What might be next?

		1) Ngram generation. (i []string, n int) where i is a token slice and n is an upper-bound for longest sequence of contiguous words wanted from ngram generation... e.g. n=3 includes 'hello' (token), 'hello world' (ngram n=2), 'hello world record' (ngram n=3)
		2) 'stop word' removal
		3) Loading result into a reverse index
	*/
	return nil
}

func tokenize(i string) ([]string, error) {
	i = strings.ToLower(i)
	tR, err := regexp.Compile("[^a-z0-9_'-]+") // perhaps we should iterate through with a more specific ruleset.
	if err != nil {
		return nil, errors.New("failed to compile the tokenizer regular expression")
	}

	return tR.Split(i, -1), nil
}
