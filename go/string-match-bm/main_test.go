package main

import (
	"strings"
	"testing"
)

func TestBf(t *testing.T) {
	cases := [][]string{
		{"abcabc", "abc"},
		{"intersting", "erst"},
		{"intersting", "in"},
		{"beautiful", "autiful"},
		{"something", "thin"},
		{"hahahah", "test"},
		{"hahah", "hahahahahah"},
	}
	// init expected index
	for _, c := range cases {
		expected := strings.Index(c[0], c[1])
		got := bf(c[0], len(c[0]), c[1], len(c[1]))
		if got != expected {
			t.Errorf("%s / %s, expected %d, got %d", c[0], c[1], expected, got)
		}
	}
}
