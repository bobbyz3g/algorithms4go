package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort_Int(t *testing.T) {
	var a = []int{6, 5, 4, 3, 2, 1}

	BubbleSort(a)

	for i := 1; i < 7; i++ {
		if a[i-1] != i {
			t.Errorf("Sort Error: should %d in index %d, not %d", i, i, a[i-1])
		}
	}
}

func TestBubbleSort_Float64(t *testing.T) {
	var tests = []struct {
		a    []float64
		want []float64
	}{
		{
			a:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
			want: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
		},
		{
			a:    []float64{1.1},
			want: []float64{1.1},
		},
		{
			a:    []float64{5.0, 4.0, 3.0, 2.0, 1.0},
			want: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
	}

	for i, tt := range tests {
		BubbleSort(tt.a)
		assert.Equal(t, tt.want, tt.a, "case %d", i)
	}
}

func TestBubbleSort_Float32(t *testing.T) {
	var tests = []struct {
		a    []float32
		want []float32
	}{
		{
			a:    []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
			want: []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
		},
		{
			a:    []float32{1.1},
			want: []float32{1.1},
		},
		{
			a:    []float32{5.0, 4.0, 3.0, 2.0, 1.0},
			want: []float32{1.0, 2.0, 3.0, 4.0, 5.0},
		},
	}

	for i, tt := range tests {
		BubbleSort(tt.a)
		assert.Equal(t, tt.want, tt.a, "case %d", i)
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
	var tests = []struct {
		a    []uint
		want []uint
	}{
		{
			a:    []uint{4, 5, 6, 3, 2, 1},
			want: []uint{1, 2, 3, 4, 5, 6},
		},
		{
			a:    []uint{1},
			want: []uint{1},
		},
		{
			a:    []uint{},
			want: []uint{},
		},
	}
	for i, tt := range tests {
		CountingSort(tt.a)
		assert.Equal(t, tt.want, tt.a, "case %d", i)
	}
}
