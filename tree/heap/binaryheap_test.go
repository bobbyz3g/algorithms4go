package heap

import (
	"github.com/Kaiser925/algorithms4go/base"
	"testing"
)

func TestMaxHeap_Push(t *testing.T) {
	heap := New(base.IntCompareFunc)
	for i := 1; i < 30; i++ {
		heap.Push(i)
	}

	if top := heap.Peek(); top != 29 {
		t.Errorf("Peek error: excepted %v, got %v", 29, top)
	}

	if capacity := cap(heap.keys); capacity != 40 {
		t.Errorf("Size grow error: excepted %v, got %v", 40, capacity)
	}
	if size := heap.Size(); size != 29 {
		t.Errorf("Size error: excepted %v, got %v", 29, size)
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	heap := New(base.IntCompareFunc)
	for i := 1; i < 30; i++ {
		heap.Push(i)
	}

	for i := 29; i > 0; i-- {
		if top := heap.Pop(); top != i {
			t.Errorf("Pop error: excepted %v, got %v", i, top)
		}
	}

	if top := heap.Pop(); top != nil {
		t.Errorf("Pop error: excepted %v, got %v", nil, top)
	}
	if capacity := cap(heap.keys); capacity != 10 {
		t.Errorf("Size shrink error, excepted %v, got %v", 10, capacity)
	}

	if size := heap.Size(); size != 0 {
		t.Errorf("Size error: excepted %v, got %v", 0, size)
	}
}
