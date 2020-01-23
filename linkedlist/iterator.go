package linkedlist

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

func (iter *Iterator) Value() interface{} {
	item := iter.current.Value
	iter.current = iter.current.next
	return item
}
