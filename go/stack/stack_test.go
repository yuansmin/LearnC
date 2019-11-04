package stack

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	cases := map[string]bool{
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"":       true,
		"([)]":   false,
		"{[]}":   true,
	}

	for s, result := range cases {
		if isValidParentheses(s) != result {
			t.Errorf("Wrong! %s != %v", s, result)
		}
	}
}
