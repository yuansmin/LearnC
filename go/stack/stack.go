package stack

type Stack struct {
	data []rune
	head int
}

func (s *Stack) Push(r rune) {
	// fmt.Printf("push %s\n", string(r))
	s.data = append(s.data, r)
	s.head++
}

func (s *Stack) Pop() rune {
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
	return &Stack{}
}

var ParentMap = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func isValidParentheses(s string) bool {
	if len(s) <= 0 {
		return true
	}
	stack := newStack()
	for _, r := range s {
		switch r {
		case '{', '(', '[':
			{
				stack.Push(r)
			}
		case '}', ')', ']':
			{
				rp, _ := ParentMap[r]
				if stack.IsEmpty() || stack.Pop() != rp {
					return false
				}
			}
		}
	}
	if !stack.IsEmpty() {
		return false
	}
	return true
}
