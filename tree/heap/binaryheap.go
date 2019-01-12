package heap

import "github.com/Kaiser925/algorithms4go/base"

type MaxHeap struct {
	keys       []interface{}
	Comparator base.CompareFunc
	size       int
}

func New(Comparator base.CompareFunc) *MaxHeap {
	return &MaxHeap{
		make([]interface{}, 20),
		Comparator,
		0,
	}
}

// Push adds a value onto the heap.
func (h *MaxHeap) Push(key interface{}) {
	capacity := cap(h.keys)
	if h.size >= capacity-1 {
		h.keys = resize(h.keys, capacity<<1)
	}
	h.size++
	h.keys[h.size] = key
	h.swim(h.size)
}

// Pop removes the max (top) element and returns it, or nil if heap is empty.
func (h *MaxHeap) Pop() interface{} {
	if h.Empty() {
		return nil
	}
	top := h.keys[1]
	h.exch(1, h.size)
	h.keys[h.size] = nil
	h.size--
	h.sink(1)

	capacity := cap(h.keys)
	if h.size < capacity>>2 {
		h.keys = resize(h.keys, capacity>>1)
	}
	return top
}

// Peek returns the top element of heap, or nil if heap is empty.
func (h *MaxHeap) Peek() interface{} {
	return h.keys[1]
}

// Size returns number of elements within the heap.
func (h *MaxHeap) Size() int {
	return h.size
}

// Empty returns true if heap does not contain any elements.
func (h *MaxHeap) Empty() bool {
	return h.size == 0
}

func (h *MaxHeap) less(i, j int) bool {
	return h.Comparator(h.keys[i], h.keys[j]) < 0
}

// Exchange elements of subscripts i,j
func (h *MaxHeap) exch(i, j int) {
	h.keys[i], h.keys[j] = h.keys[j], h.keys[i]
}

func (h *MaxHeap) swim(k int) {
	for k > 1 && h.less(k>>1, k) {
		h.exch(k>>1, k)
		k = k >> 1
	}
}

func (h *MaxHeap) sink(k int) {
	for k<<1 <= h.size {
		j := k << 1
		if j < h.size && h.less(j, j+1) {
			j++
		}
		if !h.less(k, j) {
			break
		}
		h.exch(k, j)
		k = j
	}
}

// Resize the slice of heap.
func resize(old []interface{}, cap int) []interface{} {
	if cap < 10 {
		return old
	}
	newSlice := make([]interface{}, cap, cap)
	copy(newSlice, old)
	return newSlice
}
