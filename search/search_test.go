package search

import (
	"testing"
)

var a = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}

func TestBinarySearch_Int(t *testing.T) {
	var tests = []struct {
		s    []int
		v    int
		want int
	}{
		{
			s:    []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    8,
			want: 5,
		},
		{
			s:    []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    1,
			want: 0,
		},
		{
			s:    []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    18,
			want: 9,
		},
		{
			s:    []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    0,
			want: -1,
		},
		{
			s:    []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    20,
			want: -1,
		},
	}

	for i, tt := range tests {
		got := BinarySearch(tt.s, tt.v)
		if got != tt.want {
			t.Errorf("%d. BinarySearch(%v, %v) = %v, want %v", i, tt.s, tt.v, got, tt.want)
		}
	}
}

func TestBinarySearch_Float(t *testing.T) {
	var tests = []struct {
		s    []float64
		v    float64
		want int
	}{
		{
			s:    []float64{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    8,
			want: 5,
		},
		{
			s:    []float64{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    1,
			want: 0,
		},
		{
			s:    []float64{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    18,
			want: 9,
		},
		{
			s:    []float64{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    0,
			want: -1,
		},
		{
			s:    []float64{1, 3, 4, 5, 6, 8, 8, 8, 11, 18},
			v:    20,
			want: -1,
		},
	}

	for i, tt := range tests {
		got := BinarySearch(tt.s, tt.v)
		if got != tt.want {
			t.Errorf("%d. BinarySearch(%v, %v) = %v, want %v", i, tt.s, tt.v, got, tt.want)
		}
	}
}

func TestBinarySearchLast(t *testing.T) {
	index := BinarySearchLast(a, 8)
	if index != 7 {
		t.Errorf("Search error: should 7, not %d", index)
	}
}
