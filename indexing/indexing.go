package indexing

import (
	"errors"
	"regexp"
	"strings"
)

func Tokenize(i string) ([]string, error) {
	i = strings.ToLower(i)
	tR, err := regexp.Compile("[^a-z0-9_'-]+") // perhaps we should iterate through with a more specific ruleset.
	if err != nil {
		return nil, errors.New("failed to compile the tokenizer regular expression")
	}

	return tR.Split(i, -1), nil
}
