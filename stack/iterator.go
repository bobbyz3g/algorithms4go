package stack

import "github.com/Kaiser925/algorithms4go/base"

// Iterator implements iterator for linkedlist.
type Iterator struct {
	current *node
}

func (s *Stack) Iter() *Iterator {
	return &Iterator{
		current: s.top,
	}
}

func (iter *Iterator) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator) Value() (interface{}, error) {
	if iter.current == nil {
		return nil, base.NoMoreValue
	}
	now := iter.current
	iter.current = iter.current.pre
	return now.value, nil
}
