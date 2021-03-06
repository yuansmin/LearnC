package bsearch

func FindLastLess(a []int, n, value int) int {
	low, high := 0, n-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		// fmt.Printf("low: %d, %d, %d\n", low, mid, high)
		if a[mid] <= value {
			if mid == n-1 || a[mid+1] > value {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
