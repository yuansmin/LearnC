package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func genNoHeadList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	pre := head
	for _, val := range vals[1:] {
		node := &ListNode{Val: val}
		pre.Next = node
		pre = node
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	// iteratively
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
		if cur == nil {
			break
		}
	}
	return pre
}

func reverseListR(head *ListNode) *ListNode {
	// reverse recuresively
	if head == nil {
		return head
	}
	_, newHead := r1(head, nil)
	return newHead
}

func r(node, pre *ListNode) (*ListNode, *ListNode) {
	if node == nil {
		return nil, pre
	}
	next := node.Next
	node.Next = pre
	pre = node
	return r(next, pre)
}

func r1(node, pre *ListNode) (*ListNode, *ListNode) {
	if node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		return r(next, pre)
	}
	return nil, pre
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
