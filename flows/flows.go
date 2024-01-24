package flows

import (
	"fmt"

	"github.com/elsif-maj/umbraSearch/db"
	"github.com/elsif-maj/umbraSearch/indexing"
	"github.com/elsif-maj/umbraSearch/kvstore"
	"github.com/jackc/pgx/v5"
)

type Server interface {
	GetDBConn() *pgx.Conn
	GetKVStore() kvstore.KVStore
}

// Still need to figure out where stop-word removal is going to happen
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

	// Test access to key-value store
	kvstore := server.GetKVStore()
	kvstore.Set("testKey", "testValue")

	// fmt.Println("Word Tokens: ", i)
	// fmt.Println("Word Tokens AND ngrams: ", tns)
	return nil
}
