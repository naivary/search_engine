package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/go-cmp/cmp"
)

func TestSearcher_Search(t *testing.T) {
	tests := []struct {
		query string
		want  []string
	}{
		{"pig", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"PiG", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"black cat", []string{"🐈\u200d⬛"}},
		{"foo bar baz", []string{}},
	}

	for _, tc := range tests {
		for _, runner := range weavertest.AllRunners() {
			runner.Name = fmt.Sprintf("%s/%q", runner.Name, tc.query)
			runner.Test(t, func(t *testing.T, searcher Searcher) {
				got, err := searcher.Search(context.Background(), tc.query)
				if err != nil {
					t.Fatalf("Search: %v", err)
				}
				if diff := cmp.Diff(tc.want, got); diff != "" {
					t.Fatalf("Search (-want,+got):\n%s", diff)
				}
			})
		}
	}
}
