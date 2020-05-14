package linkedlist

import "github.com/Kaiser925/algorithms4go/base"

// Iterator implements iterator for linkedlist.
type Iterator struct {
	current *Element
}

func (l *SingleLinkedList) Iter() *Iterator {
	return &Iterator{
		l.head.next,
	}
}

func (iter *Iterator) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator) Value() (interface{}, error) {
	if iter.current == nil {
		return nil, base.NoMoreValue
	}
	item := iter.current.Value
	iter.current = iter.current.next
	return item, nil
}
