package bag

type Iterator struct {
	current *node
}

// Iterator returns a new iterator of bag.
func (b *Bag) Iterator() Iterator {
	return Iterator{current: b.first}
}

// Value returns current value of iterator.
func (i *Iterator) Value() interface{} {
	item := i.current.item
	i.current = i.current.next
	return item
}

// HasNext returns true if iterator has next value, or false if iterator has not,
func (i *Iterator) HasNext() bool {
	return i.current != nil
}
