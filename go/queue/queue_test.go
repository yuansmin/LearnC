package queue

import (
	"fmt"
	"testing"
)

func TestRecentCounter(t *testing.T) {
	tt := [][][]int{
		{{4}, {10, 11, 12, 15}},
		{{3}, {10, 11, 12, 3011}},
		{{2}, {10, 12, 3011}},
		{{1501}, {}},
		{{1}, {10, 12, 3050}},
	}
	tmp := make([]int, 1504)
	// fmt.Printf("%v %v", tt, tmp)
	for i := 0; i < 1504; i++ {
		tmp[i] = i * 2
	}
	tt[3][1] = tmp
	// max := int(1E9)
	for index, l1 := range tt {
		counter := Constructor()
		var count int
		for _, ct := range l1[1] {
			count = counter.Ping(ct)
		}
		fmt.Printf("%d\n", count)
		if count != l1[0][0] {
			t.Errorf("index: %d, expected: %d, got: %d", index, l1[0][0], count)
		}
	}
}
