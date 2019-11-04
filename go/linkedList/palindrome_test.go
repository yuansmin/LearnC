package main

import (
	"testing"
)

func TestPalindrome1(t *testing.T) {
	cases := map[string]bool{
		"abcdedcba": true,
		"fuck":      false,
		"shutup":    false,
		"a":         true,
		"aa":        true,
		"aaa":       true,
		"aba":       true,
		"aaaa":      true,
		"":          false,
	}
	for k, v := range cases {
		l := genLinkedList(k)
		res := isPalindrome(l)
		if res != v {
			t.Errorf("%s: %v", k, res)
		}
	}
}
