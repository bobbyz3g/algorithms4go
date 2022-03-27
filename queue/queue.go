package queue

type node[T any] struct {
	value T
	next  *node[T]
}

type Queue[T any] struct {
	front *node[T]
	rear  *node[T]
	len   int
}

// NewQueue returns a queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		nil,
		nil,
		0,
	}
}

func (q *Queue[T]) init(val T) {
	n := &node[T]{
		value: val,
		next:  nil,
	}
	q.front = n
	q.rear = n
	q.len = 1
}

// Dequeue returns the front of queue, or nil if queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.len == 0 {
		var noop T
		return noop, false
	}
	elem := q.front
	val := elem.value
	q.front = q.front.next
	elem = nil
	q.len--
	return val, true
}

// Enqueue adds the element to queue.
func (q *Queue[T]) Enqueue(val T) {
	if q.len == 0 {
		q.init(val)
		return
	}
	n := &node[T]{
		val,
		nil,
	}
	q.rear.next = n
	q.rear = n
	q.len++
}

// Len returns length of queue.
func (q *Queue[T]) Len() int {
	return q.len
}

// Empty return true if queue was empty, or false if queue was not empty.
func (q *Queue[T]) Empty() bool {
	return q.len == 0
}
