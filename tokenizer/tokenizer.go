package tokenizer

import (
	"errors"
	"regexp"
	"strings"
)

func Tokenize(i string) ([]string, error) {
	i = strings.ToLower(i)

	tR, err := regexp.Compile("[^a-z0-9_'-]+")
	if err != nil {
		return nil, errors.New("failed to compile the tokenizer regular expression")
	}

	return tR.Split(i, -1), nil
}
