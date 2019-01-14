package circularqueue

import "fmt"

type QueueError struct {
	msg string
}

func (qe *QueueError) Error() string {
	return fmt.Sprintf("QueueError: %s", qe.msg)
}

type CircularQueue struct {
	items []interface{}
	len   int
	cap   int
	tail  int
	head  int
}

func New(cap int) *CircularQueue {
	return &CircularQueue{
		items: make([]interface{}, cap),
		len:   0,
		cap:   cap,
		tail:  0,
		head:  0,
	}
}

// Put adds the item to queue.
func (q *CircularQueue) Put(item interface{}) error {
	if q.len == q.cap {
		return &QueueError{"Queue is full"}
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.cap
	q.len++
	return nil
}

// Get retires item from the queue.
func (q *CircularQueue) Get() interface{} {
	if q.len == 0 {
		return nil
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.cap
	q.len--
	return item
}
