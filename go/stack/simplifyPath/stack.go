// https://leetcode.com/problems/simplify-path/
package stack

import (
	"fmt"
)

type Stack struct {
	data []string
	head int
}

func (s *Stack) Push(r string) {
	// fmt.Printf("push %s\n", string(r))
	s.data = append(s.data, r)
	s.head++
}

func (s *Stack) Pop() string {
	s.head--
	res := s.data[s.head]
	// fmt.Printf("pop %s\n", string(res))
	s.data = s.data[:len(s.data)-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	return s.head == 0
}

func newStack() *Stack {
	return &Stack{
		data: make([]string, 0, 10),
	}
}

// "/a/../../b/../c//.//"
func simplifyPath(path string) string {
	length := len(path)
	if length <= 1 {
		return path
	}

	stack := newStack()
	for i := 1; i < length; i++ {
		if path[i] == '/' {
			continue
		}

		// find the identifier between /.*/
		start := i
		for i+1 < length && path[i+1] != '/' {
			i++
		}
		identifier := path[start : i+1]
		switch identifier {
		case ".":
		case "..":
			{
				if !stack.IsEmpty() {
					stack.Pop()
				}
			}
		default:
			{
				stack.Push(identifier)
			}
		}
	}

	cPath := ""
	for !stack.IsEmpty() {
		cPath = fmt.Sprintf("/%s%s", stack.Pop(), cPath)
	}
	if cPath == "" {
		return "/"
	}
	return cPath
}
