package linkedlist

import (
	"testing"
)

// Test the Add with single linked list.
func TestSingleLinkedList_Add(t *testing.T) {
	l := New()
	num := 3

	for i := 0; i < num; i++ {
		l.Add(i)
	}

	if l.len != 3 {
		t.Errorf("List len error: Expected: %v, Actual: %v", num, l.len)
	}

	e := l.Front()
	for i := 0; i < num; i++ {
		if e.Value != i {
			t.Errorf("List element value error: Expected: %v, Actual: %v", i, e.Value)
		}
		e = e.Next()
	}
}

// Test Remove with single linked list.
func TestSingleLinkedList_Remove(t *testing.T) {
	l1 := New()
	num := 3

	for i := 0; i < num; i++ {
		l1.Add(i)
	}

	v := l1.Remove(2)
	if v != 2 {
		t.Errorf("List remove error: expected %v, actual %v", 3, v)
	}

	if l1.Len() != 2 {
		t.Errorf("List remove error, len of list not reduce")
	}

	v = l1.Remove(4)

	if v != nil {
		t.Errorf("List remove error, can remove the element is not in list")
	}

	l2 := New()
	v = l2.Remove(1)
	if v != nil {
		t.Errorf("List remove error, can remove the element is not in list")
	}
}

// Test Back and Front
func TestSingleLinkedList_Back_And_Front(t *testing.T) {
	l := New()
	l.Add("Front")
	l.Add("Middle")
	l.Add("Back")

	if l.Front().Value != "Front" || l.Back().Value != "Back" {
		t.Errorf("Front or Back error")
	}
}

func TestSingleLinkedList_Iterator(t *testing.T) {
	l := New()
	l.Add("a")
	l.Add("b")
	l.Add("c")

	j := 1
	iter := l.Iterator()
	for iter.HasNext() {
		val := iter.Value()
		switch j {
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
		j++
	}
}
