package stack

import "github.com/Kaiser925/algorithms4go/base"

// Iterator implements iterator for linkedlist.
type Iterator[T any] struct {
	current *node[T]
}

func (s *Stack[T]) Iter() *Iterator[T] {
	return &Iterator[T]{
		current: s.top,
	}
}

func (iter *Iterator[T]) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator[T]) Value() (T, error) {
	if iter.current == nil {
		var noop T
		return noop, base.NoMoreValue
	}
	now := iter.current
	iter.current = iter.current.pre
	return now.value, nil
}
