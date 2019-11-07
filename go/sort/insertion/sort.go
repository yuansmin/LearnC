package sort

func Sort(a []int, n int) {
	for i := -1; i < n-2; i++ {
		// find the most small element
		small := i + 1
		for j := small + 1; j < n; j++ {
			if a[j] < a[small] {
				small = j
			}
		}
		// insert to the sorted array
		j := i
		tmp := a[j+1]
		for ; j > 0; j-- {
			if a[j] > a[small] {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = a[small]
		a[small] = tmp
	}
}
