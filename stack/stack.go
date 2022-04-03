package stack

type Stack[T any] struct {
	top *node[T]
	len int
}

type node[T any] struct {
	value T
	pre   *node[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		nil,
		0,
	}
}

// Push push a value onto top of stack.
func (s *Stack[T]) Push(value T) {
	newNode := &node[T]{value, s.top}
	s.top = newNode
	s.len++
}

// Pop pops the top item of stack and returns it, or nil if stack is empty.
func (s *Stack[T]) Pop() T {
	if s.len == 0 {
		var noop T
		return noop
	}
	popNode := s.top
	s.top = popNode.pre
	s.len--
	return popNode.value
}

// Len returns the number of items in stack.
func (s *Stack[T]) Len() int {
	return s.len
}

// Peek views top item of stack, or nil if stack is empty.
func (s *Stack[T]) Peek() T {
	if s.len == 0 {
		var noop T
		return noop
	}
	return s.top.value
}
