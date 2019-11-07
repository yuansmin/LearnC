// time complexity: smallest O(n^2), avg: O(n^2), biggest: O(n^2)
package sort

func Sort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := -1; i < n-2; i++ {
		// find the smallest element
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
