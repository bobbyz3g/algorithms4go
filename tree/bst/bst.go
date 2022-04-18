package bst

import (
	"github.com/Kaiser925/algorithms4go/base"
)

type node[K any, V any] struct {
	key   K
	value V
	left  *node[K, V]
	right *node[K, V]
	n     int //The total number of nodes in the subtree Rooted at this node.
}

// BST is a binary search tree.
type BST[K, V any] struct {
	Root       *node[K, V]
	Comparator base.CompareFunc
}

// NewBST returns a new BST with Comparator.
func NewBST[K, V any](Comparator base.CompareFunc) *BST[K, V] {
	return &BST[K, V]{
		nil,
		Comparator,
	}
}

func (x *node[K, V]) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

func (t *BST[K, V]) put(x *node[K, V], key K, val V) *node[K, V] {
	if x == nil {
		x = &node[K, V]{
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

func (t *BST[K, V]) get(x *node[K, V], key K) V {
	if x == nil {
		var noop V
		return noop
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
func (t *BST[K, V]) min(x *node[K, V]) *node[K, V] {
	if x.left == nil {
		return x
	}
	return t.min(x.left)
}

func (t *BST[K, V]) max(x *node[K, V]) *node[K, V] {
	if x.right == nil {
		return x
	}
	return t.max(x.right)
}

func (t *BST[K, V]) deleteMin(x *node[K, V]) *node[K, V] {
	if x.left == nil {
		return x.right
	}
	x.left = t.deleteMin(x.left)
	x.n = 1 + x.left.size() + x.right.size()
	return x
}

func (t *BST[K, V]) deleteMax(x *node[K, V]) *node[K, V] {
	if x.right == nil {
		return x.left
	}
	x.right = t.deleteMax(x.right)
	x.n = 1 + x.left.size() + x.right.size()
	return x
}

func (t *BST[K, V]) delete(x *node[K, V], key K) *node[K, V] {
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

func (t *BST[K, V]) selects(x *node[K, V], k int) K {
	if x == nil {
		var noop K
		return noop
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

func (t *BST[K, V]) rank(x *node[K, V], key K) int {
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

func (t *BST[K, V]) keysByIndex(x *node[K, V], keys []K, lo K, hi K, cur *int) {
	if x == nil {
		return
	}

	cmplo := t.Comparator(lo, x.key)
	cmphi := t.Comparator(hi, x.key)
	if cmplo < 0 {
		t.keysByIndex(x.left, keys, lo, hi, cur)
	}

	if cmplo <= 0 && cmphi >= 0 {
		keys[*cur] = x.key
		*cur++
	}

	if cmphi > 0 {
		t.keysByIndex(x.right, keys, lo, hi, cur)
	}
}

func (t *BST[K, V]) floor(x *node[K, V], key K) *node[K, V] {
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
	}
	return x
}

func (t *BST[K, V]) ceiling(x *node[K, V], key K) *node[K, V] {
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
	}

	return x
}

// Put inserts key-value into the tree.
// If there is already a "key" in the tree, Enqueue will update the value of key.
func (t *BST[K, V]) Put(key K, val V) {
	t.Root = t.put(t.Root, key, val)
}

// Get returns value of node by its key or nil if key is not found in tree.
func (t *BST[K, V]) Get(key K) V {
	return t.get(t.Root, key)
}

// TODO:
// Contains returns true if tree contains key or false if doesn't contain.
//func (t *BST[K, V]) Contains(key K) bool {
//	var noop V
//	return t.Get(key) != noop
//}

// Delete deletes node from tree by key.
func (t *BST[K, V]) Delete(key K) {
	t.Root = t.delete(t.Root, key)
}

// DeleteMin deletes min node of tree.
func (t *BST[K, V]) DeleteMin() {
	t.Root = t.deleteMin(t.Root)
}

// DeleteMax deletes max node of tree.
func (t *BST[K, V]) DeleteMax() {
	t.Root = t.deleteMax(t.Root)
}

// Min returns the min key in tree.
func (t *BST[K, V]) Min() K {
	x := t.min(t.Root)
	if x == nil {
		var noop K
		return noop
	}
	return x.key
}

// Max returns the max key in tree.
func (t *BST[K, V]) Max() K {
	x := t.max(t.Root)
	if x == nil {
		var noop K
		return noop
	}
	return x.key
}

// Select returns the key in the symbol table whose rank is k.
// This is the (k+1)st smallest key in the symbol table.
// t.Select(k) == Keys[k]
func (t *BST[K, V]) Select(k int) K {
	if k < 0 || k > t.Size() {
		var noop K
		return noop
	}
	return t.selects(t.Root, k)
}

// TODO:
// Rank returns the number of keys in the symbol table strictly less than input key.
//func (t *BST[K, V]) Rank(key K) int {
//	var noop K
//	if key == noop {
//		return -1
//	}
//	return t.rank(t.Root, key)
//}

// Size returns number of nodes in the tree.
func (t *BST[K, V]) Size() int {
	return t.Root.size()
}

// Keys returns all keys in order.
func (t *BST[K, V]) Keys() []K {
	return t.KeysByIndex(t.Min(), t.Max())
}

// KeysByIndex returns all keys between "lo" and "hi" in order.
func (t *BST[K, V]) KeysByIndex(lo, hi K) []K {
	keys := make([]K, t.Root.size())
	cur := 0
	t.keysByIndex(t.Root, keys, lo, hi, &cur)
	return keys
}

// Floor returns floor key of the input key, or nil if no floor is found.
func (t *BST[K, V]) Floor(key K) K {
	x := t.floor(t.Root, key)
	if x == nil {
		var noop K
		return noop
	}
	return x.key
}

// Ceiling returns ceiling key of the input key, or nil if no ceiling is found.
func (t *BST[K, V]) Ceiling(key K) K {
	x := t.ceiling(t.Root, key)
	if x == nil {
		var noop K
		return noop
	}
	return x.key
}

// Empty returns true if there is no node, else return false.
func (t *BST[K, V]) Empty() bool {
	return t.Root.size() == 0
}
