package main

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	cases := [][][]int{
		{{1, 2, 3, 4}, {4, 3, 2, 1}},
		{{0, 1}, {1, 0}},
		{{0}, {0}},
		{{1, 2, 3}, {3, 2, 1}},
	}
	for _, reverse := range cases {
		l := genNoHeadList(reverse[0])
		rl := reverseList(l)
		// printList(rl)
		for _, v := range reverse[1] {
			if rl == nil || rl.Val != v {
				t.Errorf("%v", reverse[0])
				break
			}
			rl = rl.Next
		}
		if rl != nil {
			t.Errorf("%v rl not end", reverse[0])
		}
	}
}

func TestRecursivelyReverseList(t *testing.T) {
	cases := [][][]int{
		{{1, 2, 3, 4}, {4, 3, 2, 1}},
		{{0, 1}, {1, 0}},
		{{0}, {0}},
		{{1, 2, 3}, {3, 2, 1}},
	}
	for _, reverse := range cases {
		l := genNoHeadList(reverse[0])
		rl := reverseListR(l)
		printList(rl)
		for _, v := range reverse[1] {
			if rl == nil || rl.Val != v {
				t.Errorf("%v", reverse[0])
				break
			}
			rl = rl.Next
		}
		if rl != nil {
			t.Errorf("%v rl not end", reverse[0])
		}
	}
}

func TestFib(t *testing.T) {
	cases := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i, v := range cases {
		if fib(i) != v {
			t.Errorf("n: %d, res: %d", i, v)
			break
		}
	}
}
