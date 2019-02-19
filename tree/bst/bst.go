package bst

import (
	"github.com/Kaiser925/algorithms4go/base"
)

type Node struct {
	key   interface{}
	value interface{}
	left  *Node
	right *Node
	n     int //The total number of nodes in the subtree Rooted at this node.
}

// BST is a binary search tree.
type BST struct {
	Root       *Node
	Comparator base.CompareFunc
}

func NewBST(Comparator base.CompareFunc) *BST {
	return &BST{
		nil,
		Comparator,
	}
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
		return t.selects(x.right, k-n-1)
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

func (t *BST) keysByIndex(x *Node, keys *[]interface{}, lo interface{}, hi interface{}, cur *int) {
	if x == nil {
		return
	}

	cmplo := t.Comparator(lo, x.key)
	cmphi := t.Comparator(hi, x.key)
	if cmplo < 0 {
		t.keysByIndex(x.left, keys, lo, hi, cur)
	}

	if cmplo <= 0 && cmphi >= 0 {
		(*keys)[*cur] = x.key
		*cur++
	}

	if cmphi > 0 {
		t.keysByIndex(x.right, keys, lo, hi, cur)
	}
}

func (t *BST) floor(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	cmp := t.Comparator(key, x.key)
	if cmp == 0 {
		return x
	} else if cmp < 0 {
		return t.floor(x.left, key)
	}

	tmp := t.floor(x.right, key)
	if tmp != nil {
		return tmp
	} else {
		return x
	}
}

func (t *BST) ceiling(x *Node, key interface{}) *Node {
	if x == nil {
		return nil
	}
	cmp := t.Comparator(key, x.key)

	if cmp == 0 {
		return x
	} else if cmp > 0 {
		return t.ceiling(x.right, key)
	}

	tmp := t.ceiling(x.left, key)
	if tmp != nil {
		return tmp
	} else {
		return x
	}
}

// Enqueue inserts key-value into the tree.
// If there is already a "key" in the tree, Enqueue will update the value of key.
func (t *BST) Put(key, val interface{}) {
	t.Root = t.put(t.Root, key, val)
}

// Dequeue returns value of node by its key or nil if key is not found in tree.
func (t *BST) Get(key interface{}) interface{} {
	return t.get(t.Root, key)
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *BST) Contains(key interface{}) bool {
	return t.Get(key) != nil
}

// Delete deletes node from tree by key.
func (t *BST) Delete(key interface{}) {
	t.Root = t.delete(t.Root, key)
}

// DeleteMin deletes min node of tree.
func (t *BST) DeleteMin() {
	t.Root = t.deleteMin(t.Root)
}

// DeleteMax deletes max node of tree.
func (t *BST) DeleteMax() {
	t.Root = t.deleteMax(t.Root)
}

// Min returns the min value in tree.
func (t *BST) Min() interface{} {
	x := t.min(t.Root)
	if x == nil {
		return nil
	}
	return x.key
}

// Max returns the max value in tree.
func (t *BST) Max() interface{} {
	x := t.max(t.Root)
	if x == nil {
		return nil
	}
	return x.key
}

// Select returns the key in the symbol table whose rank is k.
// This is the (k+1)st smallest key in the symbol table.
// t.Select(k) == Keys[k]
func (t *BST) Select(k int) interface{} {
	if k < 0 || k > t.Size() {
		return nil
	}
	return t.selects(t.Root, k)
}

// Rank returns the number of keys in the symbol table strictly less than input key.
func (t *BST) Rank(key interface{}) int {
	if key == nil {
		return -1
	}
	return t.rank(t.Root, key)
}

// Size returns number of nodes in the tree.
func (t *BST) Size() int {
	return t.Root.size()
}

// Keys returns all keys in order.
func (t *BST) Keys() []interface{} {
	return t.KeysByIndex(t.Min(), t.Max())
}

// KeysByIndex returns all keys between "lo" and "hi" in order.
func (t *BST) KeysByIndex(lo, hi interface{}) []interface{} {
	keys := make([]interface{}, t.Root.size())
	cur := 0
	t.keysByIndex(t.Root, &keys, lo, hi, &cur)
	return keys
}

// Floor returns floor key of the input key, or nil if no floor is found.
func (t *BST) Floor(key interface{}) interface{} {
	x := t.floor(t.Root, key)
	if x == nil {
		return nil
	}
	return x.key
}

// Ceiling returns ceiling key of the input key, or nil if no ceiling is found.
func (t *BST) Ceiling(key interface{}) interface{} {
	x := t.ceiling(t.Root, key)
	if x == nil {
		return nil
	}
	return x.key
}

func (t *BST) Empty() bool {
	return t.Root.size() == 0
}
