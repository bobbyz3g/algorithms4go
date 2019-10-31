package queue

type node struct {
	value interface{}
	next  *node
}

type Queue struct {
	front *node
	rear  *node
	len   int
}

// NewQueue returns a queue.
func NewQueue() *Queue {
	return &Queue{
		nil,
		nil,
		0,
	}
}

func (q *Queue) init(val interface{}) {
	n := &node{
		value: val,
		next:  nil,
	}
	q.front = n
	q.rear = n
	q.len = 1
}

// Dequeue returns the front of queue, or nil if queue is empty.
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.len == 0 {
		return nil, false
	}
	elem := q.front
	val := elem.value
	q.front = q.front.next
	elem = nil
	q.len--
	return val, true
}

// Enqueue adds the element to queue.
func (q *Queue) Enqueue(val interface{}) {
	if q.len == 0 {
		q.init(val)
		return
	}
	n := &node{
		val,
		nil,
	}
	q.rear.next = n
	q.rear = n
	q.len++
}

// Len returns length of queue.
func (q *Queue) Len() int {
	return q.len
}

// Empty return true if queue was empty, or false if queue was not empty.
func (q *Queue) Empty() bool {
	return q.len == 0
}
