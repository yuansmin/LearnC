package sort

import "fmt"

func Sort(a []int, n int) {
	if n <= 1 {
		return
	}
	// heapfy
	heapifyUp(a, n)
	// heapifyUpSimple(a, n)

	// sort
	for i := n - 1; i >= 1; i-- {
		tmp := a[0]
		a[0] = a[i]
		heapifyDown(a, 0, i)
		a[i] = tmp
	}
}

// heapify from none leaf node
// n: length of array
func heapifyUp(a []int, n int) {
	fmt.Printf("%v\n", a)
	for mid := (n - 2) / 2; mid >= 0; mid-- {
		bigger := mid
		if mid*2+2 < n && a[mid*2+2] > a[bigger] {
			bigger = mid*2 + 2
		}
		if a[mid*2+1] > a[bigger] {
			bigger = mid*2 + 1
		}
		if mid == bigger {
			continue
		}
		a[mid], a[bigger] = a[bigger], a[mid]
		heapifyDown(a, bigger, n)
	}
}

func heapifyUpSimple(a []int, n int) {
	for i := 1; i < n; i++ {
		heapify(a, i)
	}
}

func heapify(a []int, i int) {
	if i < 1 {
		return
	}
	parent := (i - 1) / 2
	for parent >= 0 && a[parent] < a[i] {
		a[parent], a[i] = a[i], a[parent]
		i = parent
		parent = (parent - 1) / 2
	}
}

func heapifyDown(a []int, i int, n int) {
	if n < 2 {
		return
	}

	child := 2*i + 1
	for child < n {
		biggerI := i
		if child+1 < n && a[child+1] > a[biggerI] {
			biggerI = child + 1
		}
		if a[child] > a[biggerI] {
			biggerI = child
		}
		if biggerI == i {
			return
		}
		a[i], a[biggerI] = a[biggerI], a[i]
		i = biggerI
		child = 2*i + 1
	}
}
