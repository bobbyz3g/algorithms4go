package queue

import (
	"github.com/Kaiser925/algorithms4go/base"
)

// Iterator implements iterator for queue.
type Iterator[T any] struct {
	current *node[T]
}

func (q *Queue[T]) Iter() *Iterator[T] {
	return &Iterator[T]{
		q.front,
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
	item := iter.current.value
	iter.current = iter.current.next
	return item, nil
}
