package queue

import "testing"

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)

	if q.Dequeue() != nil {
		t.Errorf("Dequeue Error: Empty queue shold get nil.")
	}

	q.Enqueue(1)
	var item interface{}
	item = q.Dequeue()

	if item != 1 {
		t.Errorf("Dequeue Error: Queue should get %d, but get %d", 1, item)
	}

	for i := 1; i < 6; i++ {
		q.Enqueue(i)
	}

	err := q.Enqueue(7)

	if err == nil {
		t.Error("Enqueue Error: Enqueue item to fulled queue")
	}

	for i := 1; i < 6; i++ {
		item = q.Dequeue()
		if item != i {
			t.Errorf("Dequeue Error: Queue should get %d, but get %d", 1, item)
		}
	}
}
