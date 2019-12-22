package bsearch

import (
	"testing"
)

func TestFindFirst(t *testing.T) {
	cases := [][][]int{
		{{81, 15}, {0, 6, 11, 11, 18, 25, 28, 37, 40, 45, 47, 56, 59, 62, 74, 81, 81, 87, 89, 94}},
		{{31, 5}, {8, 15, 26, 28, 29, 31, 31, 37, 41, 47, 47, 56, 58, 66, 85, 87, 87, 88, 90, 95}},
		{{53, 9}, {0, 3, 5, 13, 21, 24, 33, 38, 47, 53, 53, 57, 59, 63, 78, 88, 89, 90, 94, 99}},
		{{2, 0}, {2, 2, 5, 10, 18, 20, 28, 46, 47, 51, 56, 61, 63, 63, 66, 76, 77, 83, 94, 96}},
		{{33, 5}, {2, 3, 5, 7, 23, 33, 33, 36, 37, 40, 41, 43, 43, 46, 52, 53, 59, 78, 91, 98}},
	}
	for i, a := range cases {
		got := FindFirst(a[1], len(a[1]), a[0][0])
		if got != a[0][1] {
			t.Errorf("%d got: %d != %d", i, got, a[0][1])
		}
	}
}