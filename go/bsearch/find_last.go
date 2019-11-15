package bsearch

func FindLast(a []int, n, value int) int {
	low, high := 0, n-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		// fmt.Printf("%d, ", mid)
		if a[mid] > value {
			high = mid - 1
		} else if a[mid] < value {
			low = mid + 1
		} else {
			if mid == n-1 || a[mid+1] != value {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
