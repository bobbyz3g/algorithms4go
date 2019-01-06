// Package linkedlist is a simple single linked list implementation in go.
// This implementations is for test purposes and pretends to demonstrate how to write in go.
package linkedlist

// SingleLinkedList represents a single linked list.
// The zero value is an empty list.
type SingleLinkedList struct {
	head Element // sentinel list element, only &head, head.next are used.
	tail *Element
	len  uint
}

// Element is a element in single linked list.
type Element struct {
	next  *Element
	Value interface{}
}

// Init initializes or clear list.
func (l *SingleLinkedList) Init() *SingleLinkedList {
	l.len = 0
	return l
}

// New returns a initialized single linked list.
func New() *SingleLinkedList {
	return new(SingleLinkedList).Init()
}

// Len returns the number of element of single linked list l.
func (l *SingleLinkedList) Len() uint {
	return l.len
}

// Add receives a new node reference to add to the current list.
func (l *SingleLinkedList) Add(v interface{}) {
	e := &Element{next: nil, Value: v}
	if l.head.next == nil {
		l.head.next = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = e
	}
	l.len++
}

// Remove removes the first element tha value is v, and returns value.
// Remove returns nil if there is no this element.
func (l *SingleLinkedList) Remove(v interface{}) interface{} {
	beforeElem := l.before(v)
	if beforeElem != nil {
		e := beforeElem.next
		beforeElem.next = e.next
		e.next = nil
		l.len--
		return e.Value
	}
	return nil
}

// Front returns the first element of list or nil if the list is empty.
func (l *SingleLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Back returns the last element of list or nil if the list is empty.
func (l *SingleLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

// before return the element before element value that is v or nil.
func (l *SingleLinkedList) before(v interface{}) *Element {
	for e := &l.head; e != nil; e = e.next {
		if e.next != nil && e.next.Value == v {
			return e
		}
	}
	return nil
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}
