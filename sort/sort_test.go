package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var a = []int{6, 5, 4, 3, 2, 1}

	BubbleSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestInsertionSort(t *testing.T) {
	var a = []int{6, 5, 4, 3, 2, 1}

	InsertionSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var a = []int{4, 5, 6, 3, 2, 1}

	SelectionSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestMergeSort(t *testing.T) {
	var a = []int{4, 5, 6, 3, 2, 1}

	MergeSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestQuickSort(t *testing.T) {
	var a = []int{4, 5, 6, 3, 2, 1}

	QuickSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestCountingSort(t *testing.T) {
	var a = []int{4, 5, 6, 3, 2, 1}

	CountingSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}
