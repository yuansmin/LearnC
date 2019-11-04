package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func run_test() {
	l1 := init_link([]int{2, 4, 3})
	// print_link(l1)
	l2 := init_link([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	print_link(l3)
}

func print_link(l1 *ListNode) {
	cur := l1
	for {
		if cur == nil {
			break
		}

		fmt.Printf("%d ->", cur.Val)
		cur = cur.Next
	}
}

func init_link(l1 []int) *ListNode {
	var first, pre *ListNode
	for i := 0; i < len(l1); i++ {
		cur := &ListNode{l1[i], nil}
		if first == nil {
			first = cur
		}
		if pre != nil {
			pre.Next = cur
		}
		pre = cur
	}
	return first
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, pre *ListNode
	cur1, cur2 := l1, l2
	step := 0
	for {
		if cur1 == nil && cur2 == nil && step == 0 {
			break
		}
		temp := step
		if cur1 != nil {
			temp += cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			temp += cur2.Val
			cur2 = cur2.Next
		}
		step = temp / 10
		value := temp % 10
		curL := &ListNode{value, nil}
		if l3 == nil {
			l3 = curL
		}
		if pre != nil {
			pre.Next = curL
		}
		pre = curL
	}
	return l3
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var bl1, bl2 []int
	for {
		bl1 = append(bl1, l1.Val)
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}
	for {
		bl2 = append(bl2, l2.Val)
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	length1, length2 := len(bl1), len(bl2)
	var n, ls int
	var long, short []int
	if length1 < length2 {
		n, ls = length2, length1
		long, short = bl2, bl1
	} else {
		n, ls = length1, length2
		long, short = bl1, bl2
	}
	step := 0
	var l3, pre *ListNode
	fmt.Println("n = %d", n)
	for i := 0; i < n; i++ {
		var tmp int
		if i < ls {
			tmp = long[n-i-1] + short[ls-i-1] + step
		} else {
			tmp = long[n-i-1] + step
		}
		step = tmp / 10
		fmt.Printf("%dX (%d)", tmp, i)
		cur := &ListNode{tmp % 10, nil}
		if pre != nil {
			pre.Next = cur
		}
		if l3 == nil {
			l3 = cur
		}
		pre = cur
	}
	if step > 0 {
		cur := &ListNode{step, nil}
		pre.Next = cur
	}
	return l3
}
