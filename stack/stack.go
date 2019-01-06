package stack

type Stack struct {
	top *node
	len int
}

type node struct {
	value interface{}
	pre   *node
}

func New() *Stack {
	return &Stack{
		nil,
		0,
	}
}

// Push push a value onto top of stack.
func (s *Stack) Push(value interface{}) {
	newNode := &node{value, s.top}
	s.top = newNode
	s.len++
}

// Pop pops the top item of stack and returns it.
// If stack is empty, Peek will return nil.
func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	popNode := s.top
	s.top = popNode.pre
	s.len--
	return popNode.value
}

// Len returns the number of items in stack.
func (s *Stack) Len() int {
	return s.len
}

// Peek views top item of stack.
// If stack is empty, Peek will return nil.
func (s *Stack) Peek() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.top.value
}
