// time complexity: smallest O(nlogn), avg: O(nlogn), biggest: O(nlogn)
// space complexity: O(n)
package sort

// for tmp array, avoid alloc mem multi times
var tmp []int

func Sort(a []int, n int) {
	if n <= 1 {
		return
	}
	tmp = make([]int, n)
	sortInternal(a, 0, n-1)
}

func sortInternal(a []int, start, end int) {
	// check end
	if start >= end-1 {
		if a[start] > a[end] {
			a[start], a[end] = a[end], a[start]
		}
		return
	}

	// divide
	mid := (start + end) / 2
	sortInternal(a, start, mid)
	sortInternal(a, mid+1, end)

	// merge
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	k := 0
	i, j := start, mid+1
	for i <= mid && j <= end {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = a[i]
		i++
		k++
	}

	for j <= end {
		tmp[k] = a[j]
		j++
		k++
	}
	for i := 0; start+i <= end; i++ {
		a[start+i] = tmp[i]
	}
}
