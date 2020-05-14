package bag

// Iterator implements Iterator for Bag.
type Iterator struct {
	current *node
}

// Iter returns a new iterator of bag.
func (b *Bag) Iter() *Iterator {
	return &Iterator{current: b.first}
}

// Value returns current value of iterator.
func (i *Iterator) Value() (interface{}, error) {
	if i.current == nil {
		return nil, nil
	}
	item := i.current.item
	i.current = i.current.next
	return item, nil
}

// HasNext returns true if iterator has next value, or false if iterator has not,
func (i *Iterator) HasNext() bool {
	return i.current != nil
}
