package queue

type Iterator struct {
	current *node
}

func (q *Queue) Iterator() Iterator {
	return Iterator{
		q.front,
	}
}

func (iter *Iterator) HasNext() bool {
	return iter.current != nil
}

func (iter *Iterator) Value() interface{} {
	item := iter.current.value
	iter.current = iter.current.next
	return item
}
