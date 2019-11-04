package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	isOdd := (n1 + n2) % 2
	medianI := (n1 + n2) / 2
	nums := make([]int, 0)
	i, j := 0, 0
	var median float64
	for {
		length := len(nums)
		if length >= medianI+1 {
			if isOdd > 0 {
				median = float64(nums[length-1])
			} else {
				median = (float64(nums[length-1]) + float64(nums[length-2])) / 2
			}
			break
		}
		var small int
		if i < n1 {
			small = nums1[i]
		}
		if j < n2 && small > nums2[j] {
			small = nums2[j]
			j++
		} else if i < n1 {
			small = nums1[i]
			i++
		} else if j < n2 {
			small = nums2[j]
			j++
		} else {
			break
		}
		nums = append(nums, small)
	}
	return median
}

func testFindMedianSortedArrays() {
	nums := [][][]int{
		{{1, 3}, {2}},
		{{1, 2}, {3, 4}},
		{{7, 9, 10, 11, 12, 14}, {8, 13, 17, 19}},
	}
	for i := 0; i < len(nums); i++ {
		res := findMedianSortedArrays(nums[i][0], nums[i][1])
		fmt.Printf("%v\n%v\nresult: %f\n\n", nums[i][0], nums[i][1], res)
	}
}
