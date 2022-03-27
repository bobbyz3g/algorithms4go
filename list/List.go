// Package list implements a single linked list.
package list

// List is a singly linked list.
type List[T any] struct {
	head *element[T]
	tail *element[T]
	len  int
}

// element is an element of a list.
type element[T any] struct {
	Prev  *element[T]
	Next  *element[T]
	Value T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func FromSlice[T any](s []T) *List[T] {
	l := NewList[T]()
	for _, v := range s {
		l.Push(v)
	}
	return l
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) Push(v T) {
	e := &element[T]{Value: v}
	if l.head == nil {
		l.head = e
		l.tail = e
	} else {
		e.Prev = l.tail
		l.tail.Next = e
		l.tail = e
	}
	l.len++
}

func (l *List[T]) Get(i int) (T, bool) {
	if i < 0 || i > l.len-1 {
		var noop T
		return noop, false
	}
	e := l.head
	for j := 0; j < i; j++ {
		e = e.Next
	}
	return e.Value, true
}

func (l *List[T]) Pop() (T, bool) {
	if l.len == 0 {
		var noop T
		return noop, false
	}
	e := l.tail
	if l.tail.Prev != nil {
		l.tail = l.tail.Prev
	}
	l.tail.Next = nil
	l.len--
	return e.Value, true
}

func (l *List[T]) ToSlice() []T {
	s := make([]T, l.len)
	e := l.head
	for i := 0; i < l.len; i++ {
		s[i] = e.Value
		e = e.Next
	}
	return s
}
