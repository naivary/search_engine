package main

import (
	"context"
)

type Searcher interface{}

type searcher struct{}

func (s searcher) Search(ctx context.Context, query string) ([]string, error) {
	// Perform the search with the following rules:
	//
	// 1. Lowercase the query and split it into words -> (strings.Fields, strings.ToLower)
	// For example, the query "Black cat" is tokenized to the words "black" and "cat".
	//
	// 2. An emoji matches a query if every word in the query is one of the emoji's labels. -> (slices.Contains)
	// For example, the cat emoji has labels ["animal", "cat"]. It does NOT
	// match the "Black cat" query because "black" is not a label. The black
	// cat emoji, on the other hand, has labels ["animal", "black", "cat"] and
	// therefore does match the query "Black cat".
	//
	// We iterate over all emojis and return the ones that match the query (sorted) -> (sort.Strings).
	//
	results := []string{}
	return results, nil
}

// matches returns whether words is a subset of labels.
func matches(labels, words []string) bool {
	return true
}
