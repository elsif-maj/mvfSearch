package indexing

import (
	"errors"
	"regexp"
	"strings"
)

func TokenizeWords(i string) ([]string, error) {
	i = strings.ToLower(i)
	tR, err := regexp.Compile("[^a-z0-9_'-]+") // perhaps we should iterate through with a more specific ruleset.
	if err != nil {
		return nil, errors.New("failed to compile the tokenizer regular expression")
	}

	return tR.Split(i, -1), nil
}

// (r)eference slice, (p)roduct slice, (l)imit
func MakeWordNgrams(r []string, p []string, l int) ([]string, error) {
	// Guard clause for l < 2 or l > 5(?)

	// i will represent the "round" of iteration, as well as the order/size (in word tokens) of the current round's ngrams that are being collected
	for i := 2; i <= l; i++ {
		var left int
		for right := i; right < len(r); right++ {
			ngram := strings.Join(r[left:right], " ")
			p = append(p, ngram)
			left++
		}
	}
	return p, nil
}
