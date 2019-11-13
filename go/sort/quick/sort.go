package sort

func Sort(a []int, n int) {
	if n <= 1 {
		return
	}

	quickSort(a, 0, n-1)
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	value := a[end]
	i := start
	for j := start; j < end; j++ {
		if a[j] < value {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[end] = a[end], a[i]
	quickSort(a, start, i-1)
	quickSort(a, i+1, end)
}
