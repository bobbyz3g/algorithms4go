package stack

import "testing"

func TestStack(t *testing.T) {
	s := New[int]()

	if s.Len() != 0 {
		t.Errorf("Length of empty stack should be 0.")
	}

	if s.Peek() != 0 {
		t.Errorf("Empty stack should peek nil.")
	}

	if s.Pop() != 0 {
		t.Errorf("Empty stack should pop nil.")
	}

	s.Push(1)

	if s.Len() != 1 {
		t.Errorf("Length should be 1.")
	}

	if s.Peek() != 1 {
		t.Errorf("Top item should be 1.")
	}

	if s.Pop() != 1 {
		t.Errorf("Stack should pop 1.")
	}

	s.Push(2)
	s.Push(3)

	if s.Peek() != 3 {
		t.Errorf("Top item should be 3.")
	}

	if s.Pop() != 3 {
		t.Errorf("Stack should pop 3.")
	}
}

func TestStack_Iterator(t *testing.T) {
	s := New[int]()

	for i := 10; i > 0; i-- {
		s.Push(i)
	}

	iter := s.Iter()
	j := 1
	for iter.HasNext() {
		if v, _ := iter.Value(); v != j {
			t.Errorf("Iterator Error: excepted %v, got %v", j, v)
		}
		j++
	}
}
