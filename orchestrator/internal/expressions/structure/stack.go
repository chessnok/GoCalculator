package structure

type Stack struct {
	top    *Node
	length int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() any {
	if s.length == 0 {
		return ""
	}
	return s.top.Value
}

func (s *Stack) Pop() any {
	if s.length == 0 {
		return ""
	}

	n := s.top
	s.top = n.Prev
	s.length--
	return n.Value
}

func (s *Stack) Push(value any) {
	n := &Node{value, s.top, nil}
	s.top = n
	s.length++
}
