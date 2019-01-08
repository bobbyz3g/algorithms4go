package tree

import "testing"

func TestRedBlackTree_Contain(t *testing.T) {
	tree := NewRedBlackTree()
	for i := 1; i < 200; i++ {
		tree.Put(IntCmp(i), i)
	}

	if tree.IsEmpty() {
		t.Errorf("Error")
	}

	tree.Delete(IntCmp(30))
	tree.DeleteMin()
	tree.DeleteMax()

	if tree.Get(IntCmp(1)) != nil || tree.Get(IntCmp(199)) != nil || tree.Get(IntCmp(30)) != nil {
		t.Errorf("Error Delete")
	}
}
