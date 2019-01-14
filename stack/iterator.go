package stack

type Iterator struct {
	stack   *Stack
	current *node
}

func (s *Stack) Iterator() Iterator {
	return Iterator{
		stack:   s,
		current: s.top,
	}
}

func (iter *Iterator) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator) Value() interface{} {
	now := iter.current
	iter.current = iter.current.pre
	return now.value
}
