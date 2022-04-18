package bst

import (
	"github.com/Kaiser925/algorithms4go/base"
	"testing"
)

var tests = [][]interface{}{
	{1, "1"},
	{2, "2"},
	{3, "3"},
	{5, "5"},
	{10, "10"},
	{6, "6"},
	{102, "102"},
	{99, "99"},
	{1, "a"}, // Update key 1.
}

var keys = []interface{}{1, 2, 3, 5, 6, 10, 99, 102}
var values = []interface{}{"a", "2", "3", "5", "6", "10", "99", "102"}

func TestBST_Put(t *testing.T) {
	tree := NewBST[int, string](base.IntCompareFunc)
	for _, test := range tests {
		k := test[0].(int)
		v := test[1].(string)
		tree.Put(k, v)
	}

	if act := tree.Size(); act != 8 {
		t.Errorf("Size Error: excepted %v, got %v", 8, act)
	}

	k := tree.Keys()
	for i, v := range k {
		if keys[i] != v {
			t.Errorf("Keys Error: excepted %v, got %v", keys[i], v)
		}
	}

	for i, v := range k {
		act := tree.Get(v)
		if values[i] != act {
			t.Errorf("Dequeue Error: excepted %v, got %v", values[i], act)
		}
	}

	if act := tree.Ceiling(7); act != 10 {
		t.Errorf("Ceiling Error: excepted %v, got %v", 10, act)
	}

	if act := tree.Floor(6); act != 6 {
		t.Errorf("Floor Error: excepted %v, got %v", 6, act)
	}

	if act := tree.Empty(); act == true {
		t.Errorf("Empty Error: excepted %v, got %v", false, act)
	}

	for i, v := range keys {
		if act := tree.Select(i); act != v {
			t.Errorf("Select Error: excepted %v, got %v", v, act)
		}
	}

	//for i, v := range keys {
	//	if act := tree.Rank(v); act != i {
	//		t.Errorf("Rank Error: excepted %v, got %v", i, act)
	//	}
	//}
}

func TestBST_Delete(t *testing.T) {
	tree := NewBST[int, string](base.IntCompareFunc)
	for _, test := range tests {
		tree.Put(test[0].(int), test[1].(string))
	}
	k := []interface{}{2, 3, 6, 10, 99}
	tree.DeleteMax()
	tree.DeleteMin()
	tree.Delete(5)

	for i, v := range k {
		if k[i] != v {
			t.Errorf("Delete Error")
		}
	}
}
