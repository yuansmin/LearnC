package sort

func Sort(a []int, n int) {
	if n <= 1 {
		return
	}
	// heapfy
	for i := 1; i < n; i++ {
		heapify(a, i)
	}

	// sort
	for i := n - 1; i >= 1; i-- {
		tmp := a[0]
		a[0] = a[i]
		heapifyDown(a, 0, i)
		a[i] = tmp
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
