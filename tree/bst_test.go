package tree

import "testing"

func TestIntBST(t *testing.T) {
	tree := NewIntBST()
	if tree.Get(1) != nil {
		t.Errorf("Error")
	}
	for i := 10; i > 0; i-- {
		tree.Put(i, i)
	}

	tree.Delete(10)

	if tree.Size() != 9 {
		t.Errorf("Error size, size is %v", tree.Size())
	}
	if tree.Min() != 1 {
		t.Errorf("Error min")
	}
	if tree.Max() != 9 {
		t.Errorf("Error max, max is %v", tree.Max())
	}

	a := tree.Keys()
	for i := 0; i < 9; i++ {
		if a[i] != (i + 1) {
			t.Errorf("Error Keys")
		}
	}
}
