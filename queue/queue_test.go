package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue[int]()
	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Empty error: expected %v, got %v", actualValue, true)
	}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if actualValue := queue.Empty(); actualValue != false {
		t.Errorf("Empty error: expected %v, got %v", actualValue, false)
	}

	if actualValue := queue.Len(); actualValue != 3 {
		t.Errorf("Len error: expected %v, got %v", actualValue, 3)
	}

}

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Dequeue()

	if actualValue, ok := queue.Dequeue(); actualValue != 2 || !ok {
		t.Errorf("Dequeue error: expected %v, got %v", actualValue, 2)
	}

	if actualValue, ok := queue.Dequeue(); actualValue != 3 || !ok {
		t.Errorf("Dequeue error: expected %v, got %v", actualValue, 3)
	}

	if actualValue := queue.Empty(); actualValue != true {
		t.Errorf("Empty error: expected %v, got %v", actualValue, true)
	}
}

func TestQueue_Iterator(t *testing.T) {
	queue := NewQueue[string]()
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	iter := queue.Iter()

	index := 1
	for iter.HasNext() {
		val, _ := iter.Value()
		switch index {
		case 1:
			if val != "a" {
				t.Errorf("Iterator error: excepted %v, got %v", "a", val)
			}
		case 2:
			if val != "b" {
				t.Errorf("Iterator error: excepted %v, got %v", "b", val)
			}
		case 3:
			if val != "c" {
				t.Errorf("Iterator error: excepted %v, got %v", "c", val)
			}

		default:
			t.Errorf("Iterator error: too many value.")
		}

		index++
	}
}
