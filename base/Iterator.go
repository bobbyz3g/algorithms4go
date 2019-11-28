package base

// Iterator is the interface that wraps the basic iterator methods.
type Iterator interface {
	// HasNext returns true if there was a next element, or false if there was not.
	HasNext() bool
	// Value returns the current element's value and moves the iterator to the next element.
	Value() interface{}
}
