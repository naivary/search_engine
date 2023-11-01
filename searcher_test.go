package main

import (
	"testing"
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
		_ = tc
	}
}
