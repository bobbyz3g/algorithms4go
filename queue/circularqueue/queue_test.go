package circularqueue

import "testing"

func TestCircularQueue(t *testing.T) {
	q := New(5)

	if q.Get() != nil {
		t.Errorf("Get Error: Empty queue shold get nil.")
	}

	q.Put(1)
	var item interface{}
	item = q.Get()

	if item != 1 {
		t.Errorf("Get Error: Queue should get %d, but get %d", 1, item)
	}

	for i := 1; i < 6; i++ {
		q.Put(i)
	}

	err := q.Put(7)

	if err == nil {
		t.Error("Put Error: Put item to fulled queue")
	}

	for i := 1; i < 6; i++ {
		item = q.Get()
		if item != i {
			t.Errorf("Get Error: Queue should get %d, but get %d", 1, item)
		}
	}
}
