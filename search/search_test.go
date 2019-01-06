package search

import (
	"testing"
)

var a = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}

func TestBinarySearch(t *testing.T) {
	index := BinarySearch(a, 8)
	if index != 5 {
		t.Errorf("Search error: should 5, not %d", index)
	}
}

func TestBinarySearchLast(t *testing.T) {
	index := BinarySearchLast(a, 8)
	if index != 7 {
		t.Errorf("Search error: should 7, not %d", index)
	}
}
