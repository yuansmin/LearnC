// check if a string linked list is reversed word
package main

import "fmt"

func main() {
	cases := map[string]bool{
		"abcdedcba": true,
		// "fuck":      false,
		// "shutup":    false,
		// "a":         true,
		// "aa":        true,
		// "aaa":       true,
		// "aba":       true,
		// "aaaa":      true,
	}
	success := true
	for k, v := range cases {
		l := genLinkedList(k)
		res := isPalindrome(l)
		state := "success"
		if res != v {
			success = false
			state = "failed"
		}
		fmt.Printf("%s: %s\n", k, state)
	}
	result := "OK"
	if !success {
		result = "Bad"
	}
	fmt.Printf("%s \n", result)
}

type Node struct {
	Data rune
	Next *Node
}

func genLinkedList(s string) *Node {
	head := &Node{}
	if len(s) == 0 {
		return head
	}
	parent := head
	for _, v := range s {
		node := &Node{Data: v}
		parent.Next = node
		parent = node
	}
	return head
}

func isPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}

	mid, end := head, head
	var pre *Node
	for {
		end = end.Next
		tmp := mid
		mid = mid.Next
		tmp.Next = pre
		pre = tmp
		if end == nil {
			break
		} else if end.Next == nil {
			pre = &Node{Data: mid.Data, Next: pre}
			break
		}
		end = end.Next
	}

	for pre != nil && mid != nil {
		if pre.Data != mid.Data {
			return false
		}
		pre, mid = pre.Next, mid.Next
	}
	return true
}

// func isPalindromeOld(head *Node) bool {
// 	if head == nil {
// 		return true
// 	}

// 	next, end := head, head
// 	var pre *Node
// 	for end != nil {
// 		if end.Next == nil {
// 			pre = &Node{Data: next.Data, Next: pre}
// 			break
// 		}
// 		end = end.Next.Next
// 		tmp := next
// 		next = next.Next
// 		tmp.Next = pre
// 		pre = tmp
// 	}

// 	fmt.Printf("test: pre: %s, next: %s\n", string(pre.Data), string(next.Data))

// 	for pre != nil && next != nil {
// 		if pre.Data != next.Data {
// 			return false
// 		}
// 		pre = pre.Next
// 		next = next.Next
// 	}
// 	return true
// }
