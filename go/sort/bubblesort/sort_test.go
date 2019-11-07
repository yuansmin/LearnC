package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	cases := [][]int{}
	for i := 0; i < 20; i++ {
		data := make([]int, 50)
		for i, _ := range data {
			data[i] = rand.Intn(100)
		}
		cases = append(cases, data)
	}

	for _, data := range cases {
		sorted := make([]int, 50)
		copy(sorted, data)
		sort.Ints(sorted)
		fmt.Printf("%v %v\n", data, sorted)
		Sort(data)
		fmt.Printf("sorted: %v\n", data)
		for i, _ := range data {
			if data[i] != sorted[i] {
				t.Errorf("%v != %v", data, sorted)
			}
		}
	}
}
