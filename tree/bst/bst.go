package bst

import (
	"github.com/Kaiser925/algorithms4go/base"
)

type Node struct {
	key   interface{}
	value interface{}
	left  *Node
	right *Node
	n     int //The total number of nodes in the subtree rooted at this node.
}

// BST is a binary search tree.
type BST struct {
	root       *Node
	Comparator base.CompareFunc
}

func (x *Node) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

func (t *BST) put(x *Node, key, val interface{}) *Node {
	if x == nil {
		x = &Node{
			key:   key,
			value: val,
			left:  nil,
			right: nil,
			n:     1,
		}
	}
	cmp := t.Comparator(key, x.key)
	if cmp < 0 {
		x.left = t.put(x.left, key, val)
	} else if cmp > 0 {
		x.right = t.put(x.right, key, val)
	} else {
		x.value = val
	}
	x.n = 1 + x.left.size() + x.right.size()
	return x
}

func (t *BST) get(x *Node, key interface{}) interface{} {
	if x == nil {
		return nil
	}
	cmp := t.Comparator(key, x.key)

	if cmp < 0 {
		return t.get(x.left, key)
	} else if cmp > 0 {
		return t.get(x.right, key)
	} else {
		return x.value
	}
}

// min returns the min node in tree.
func (t *BST) min(x *Node) *Node {
	if x.left == nil {
		return x
	}
	return t.min(x.left)
}

func (t *BST) max(x *Node) *Node {
	if x.right == nil {
		return x
	}
	return t.max(x.right)
}

func (t *BST) deleteMin(x *Node) *Node {
	if x.left == nil {
		return x.right
	}
	x.left = t.deleteMin(x.left)
	x.n = 1 + x.left.size() + x.right.size()
	return x
}

func (t *BST) deleteMax(x *Node) *Node {
	if x.right == nil {
		return x.left
	}
	x.right = t.deleteMax(x.right)
	x.n = 1 + x.left.size() + x.right.size()
	return x
}

func (t *BST) delete(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	cmp := t.Comparator(key, x.key)
	if cmp < 0 {
		x.left = t.delete(x.left, key)
	} else if cmp > 0 {
		x.right = t.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}

		tmp := x
		x = t.min(tmp.right)
		x.right = t.deleteMin(x.right)
		x.left = tmp.left
	}
	return x
}

func (t *BST) selects(x *Node, k int) interface{} {
	if x == nil {
		return nil
	}

	n := x.left.size()
	if n > k {
		return t.selects(x.left, k)
	} else if n < k {
		return t.selects(x.right, k)
	} else {
		return x.key
	}
}

func (t *BST) rank(x *Node, key interface{}) int {
	if x == nil {
		return 0
	}

	cmp := t.Comparator(key, x.key)
	if cmp < 0 {
		return t.rank(x.left, key)
	} else if cmp > 0 {
		return 1 + x.left.size() + t.rank(x.right, key)
	} else {
		return x.left.size()
	}
}

// Put inserts key-value into the tree.
// If there is already a "key" in the tree, Put will update the value of key.
func (t *BST) Put(key, val interface{}) {
	t.root = t.put(t.root, key, val)
}

// Get returns value of node by its key or nil if key is not found in tree.
func (t *BST) Get(key interface{}) interface{} {
	return t.get(t.root, key)
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *BST) Contains(key interface{}) bool {
	return t.Get(key) != nil
}

// Delete deletes node from tree by key.
func (t *BST) Delete(key interface{}) {
	t.root = t.delete(t.root, key)
}

// Min returns the min value in tree.
func (t *BST) Min() interface{} {
	return t.min(t.root)
}

// Max returns the max value in tree.
func (t *BST) Max() interface{} {
	return t.max(t.root)
}

// Select returns Kth key.
func (t *BST) Select(k int) interface{} {
	return t.selects(t.root, k)
}

// Rank returns key's ranking.
func (t *BST) Rank(key interface{}) int {
	return t.rank(t.root, key)
}
