package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_Get(t *testing.T) {
	l := NewList[int]()
	_, ok := l.Get(0)
	assert.False(t, ok)
	_, ok = l.Get(10)
	assert.False(t, ok)
}

func TestList_Pop(t *testing.T) {
	l := NewList[int]()
	_, ok := l.Pop()
	assert.False(t, ok)

}

func TestList_Pop2(t *testing.T) {
	tests := []struct {
		s    []int
		want int
		ok   bool
	}{
		{s: []int{}, want: 0, ok: false},
		{s: []int{1, 2, 3}, want: 3, ok: true},
	}
	for _, tt := range tests {
		l := FromSlice(tt.s)
		got, ok := l.Pop()
		assert.Equal(t, tt.want, got)
		assert.Equal(t, tt.ok, ok)
		if ok {
			assert.Equal(t, len(tt.s)-1, l.Len())
		}
	}
}

func TestFromSlice_int(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2, 3},
	}
	for _, tt := range tests {
		l := FromSlice(tt)
		assert.Equal(t, len(tt), l.Len())
		for i, v := range tt {
			got, _ := l.Get(i)
			assert.Equal(t, v, got, "index: %d want %d got %d", i, v, got)
		}
	}
}

func TestFromSlice_String(t *testing.T) {
	tests := [][]string{
		{},
		{"a"},
		{"a", "b", "c"},
	}
	for _, tt := range tests {
		l := FromSlice(tt)
		assert.Equal(t, len(tt), l.Len())
		for i, v := range tt {
			got, _ := l.Get(i)
			assert.Equal(t, v, got, "index: %d want %d got %d", i, v, got)
		}
	}
}

func TestList_ToSlice(t *testing.T) {
	tests := []struct {
		s    []int
		want []int
	}{
		{s: []int{}, want: []int{}},
		{s: []int{1}, want: []int{1}},
		{s: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		l := FromSlice(tt.s)
		got := l.ToSlice()
		assert.Equal(t, tt.want, got)
	}
}
