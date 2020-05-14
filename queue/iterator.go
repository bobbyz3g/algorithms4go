package queue

import (
	"github.com/Kaiser925/algorithms4go/base"
)

// Iterator implements iterator for linkedlist.
type Iterator struct {
	current *node
}

func (q *Queue) Iter() *Iterator {
	return &Iterator{
		q.front,
	}
}

func (iter *Iterator) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator) Value() (interface{}, error) {
	if iter.current == nil {
		return nil, base.NoMoreValue
	}
	item := iter.current.value
	iter.current = iter.current.next
	return item, nil
}
