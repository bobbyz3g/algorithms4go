package base

// Iterable is a iterator for containers.
type Iterable interface {
	// HasNext returns true if there was a next element, or false if there was not.
	HasNext() bool
	// Value returns the current element's value and moves the iterator to the next element.
	Value() interface{}
}
