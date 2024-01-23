package flows

import (
	"fmt"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/elsif-maj/umbraSearch/indexing"
	"github.com/jackc/pgx/v5"
)

type Server interface {
	GetDBConn() *pgx.Conn
}

func ProcessInputAsWords(server Server, id int) error {
	// Get snippet from database
	snippet, err := db.GetSnippet(server.GetDBConn(), id)
	if err != nil {
		return fmt.Errorf("failed to get snippet from database: %w", err)
	}

	// Tokenize words from snippet -- come back to this and add the title
	i, err := indexing.TokenizeWords(snippet.Code)
	if err != nil {
		return fmt.Errorf("failed to tokenize snippet id: %d", id)
	}

	// (t)okens and (n)gram(s) slice (tns) will be a step-by-step 'running total' slice of tokens and ngrams that is appended-to each step of the way
	// (i) will remain unchanged as a reference of the word tokens
	tns := []string(i)

	// Make word-Ngrams from word tokens
	tns, err = indexing.MakeWordNgrams(i, tns, 3)
	if err != nil {
		return fmt.Errorf("failed to tokenize snippet id: %d", id)
	}

	//
	// fmt.Println("Word Tokens: ", i)
	// fmt.Println("Word Tokens AND ngrams: ", tns)

	/*
		What might be next?

		1) Ngram generation. (i []string, n int) where i is a token slice and n is an upper-bound for longest sequence of contiguous words wanted from ngram generation... e.g. n=3 includes 'hello' (token), 'hello world' (ngram n=2), 'hello world record' (ngram n=3)
		2) 'stop word' removal
		3) Loading result into a reverse index
	*/
	return nil
}
