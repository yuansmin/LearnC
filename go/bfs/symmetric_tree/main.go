// leetcode: https://leetcode.com/problems/symmetric-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0, 20)
	next_queue := make([]*TreeNode, 0, 20)
	queue = append(queue, root)
	data := make([]*int, 0, 20)
	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			n := queue[i]
			if n.Left == nil {
				data = append(data, nil)
			} else {
				data = append(data, &n.Left.Val)
				next_queue = append(next_queue, n.Left)
			}
			if n.Right == nil {
				data = append(data, nil)
			} else {
				data = append(data, &n.Right.Val)
				next_queue = append(next_queue, n.Right)
			}
		}
		if !isSymmetricList(data, len(data)) {
			return false
		}
		queue, next_queue = next_queue, queue[:0]
		data = data[:0]
	}
	return true
}

func isSymmetricList(data []*int, n int) bool {
	for i, j := 0, n-1; i < j; i++ {
		if data[i] != nil && data[j] != nil {
			if *data[i] != *data[j] {
				return false
			}
		} else if data[i] != data[j] {
			return false
		}
		j--
	}
	return true
}
