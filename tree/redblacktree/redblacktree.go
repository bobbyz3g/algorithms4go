package redblacktree

import (
	"github.com/Kaiser925/algorithms4go/base"
)

type color bool

const (
	// RED represents whether it is a red node
	RED color = true
	// BLACK represents whether it is a black node
	BLACK color = false
)

// Node represents the node of rbtree.
type Node struct {
	key   interface{}
	value interface{}
	left  *Node
	right *Node
	n     int
	color color
}

func (x *Node) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

func (x *Node) isRed() bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

// Tree is the struct of red-black tree.
type Tree struct {
	Root       *Node
	Comparator base.CompareFunc
}

// New returns new redblacktree with compareFunc.
func New(compareFunc base.CompareFunc) *Tree {
	return &Tree{nil, compareFunc}
}

// Put inserts key-value into the tree.
// If there is already a "key" in the tree, Enqueue will update the value of key.
func (t *Tree) Put(key, val interface{}) {
	t.Root = t.put(t.Root, key, val)
	t.Root.color = BLACK
}

// Get returns value of node by its key or nil if key is not found in tree.
func (t *Tree) Get(key interface{}) interface{} {
	return t.get(t.Root, key)
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *Tree) Contains(key interface{}) bool {
	return t.Get(key) != nil
}

// Delete deletes node from tree by key.
func (t *Tree) Delete(key interface{}) {
	if key == nil {
		return
	}

	if t.Empty() {
		return
	}

	if !t.Root.left.isRed() && !t.Root.right.isRed() {
		t.Root.color = RED
	}

	t.Root = t.delete(t.Root, key)

	if !t.Empty() {
		t.Root.color = BLACK
	}
}

// DeleteMin deletes the min node.
func (t *Tree) DeleteMin() {
	if t.Empty() {
		return
	}
	if !t.Root.left.isRed() && !t.Root.right.isRed() {
		t.Root.color = RED
	}

	t.Root = t.deleteMin(t.Root)

	if !t.Empty() {
		t.Root.color = BLACK
	}
}

// DeleteMax deletes the max node.
func (t *Tree) DeleteMax() {
	if t.Empty() {
		return
	}
	if !t.Root.left.isRed() && !t.Root.right.isRed() {
		t.Root.color = RED
	}

	t.Root = t.deleteMax(t.Root)

	if !t.Empty() {
		t.Root.color = BLACK
	}
}

// Ceiling returns ceiling key of the input key, or nil if no ceiling is found.
func (t *Tree) Ceiling(key interface{}) interface{} {
	x := t.ceiling(t.Root, key)
	if x == nil {
		return nil
	}
	return x.key
}

// Floor returns floor key of the input key, or nil if no floor is found.
func (t *Tree) Floor(key interface{}) interface{} {
	x := t.floor(t.Root, key)
	if x == nil {
		return nil
	}
	return x.key
}

// Min returns the min value in tree.
func (t *Tree) Min() interface{} {
	x := t.min(t.Root)
	if x == nil {
		return nil
	}
	return x.key
}

// Max returns the max value in tree.
func (t *Tree) Max() interface{} {
	x := t.max(t.Root)
	if x == nil {
		return nil
	}
	return x.key
}

// Select returns the key in the symbol table whose rank is k.
// This is the (k+1)st smallest key in the symbol table.
// t.Select(k) == Keys[k]
func (t *Tree) Select(k int) interface{} {
	if k < 0 || k > t.Size() {
		return nil
	}
	return t.selects(t.Root, k)
}

// Rank returns the number of keys in the symbol table strictly less than input key.
func (t *Tree) Rank(key interface{}) int {
	if key == nil {
		return -1
	}
	return t.rank(t.Root, key)
}

// Size returns number of nodes in the tree.
func (t *Tree) Size() int {
	return t.Root.size()
}

// Keys returns all keys in order.
func (t *Tree) Keys() []interface{} {
	return t.KeysByIndex(t.Min(), t.Max())
}

// KeysByIndex returns all keys between "lo" and "hi" in order.
func (t *Tree) KeysByIndex(lo, hi interface{}) []interface{} {
	keys := make([]interface{}, t.Root.size())
	cur := 0
	t.keysByIndex(t.Root, &keys, lo, hi, &cur)
	return keys
}

// Empty returns tree is empty or not.
func (t *Tree) Empty() bool {
	return t.Root.size() == 0
}

func (t *Tree) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (t *Tree) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (t *Tree) flipColors(h *Node) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// moveRedLeft assuming that h is red and both h.left and h.left.left
// are black, make h.left or one of its children red.
func (t *Tree) moveRedLeft(h *Node) *Node {
	t.flipColors(h)
	if h.right.left.isRed() {
		h.right = t.rotateRight(h.right)
		h = t.rotateLeft(h)
		t.flipColors(h)
	}
	return h
}

// moveRedRight assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func (t *Tree) moveRedRight(h *Node) *Node {
	t.flipColors(h)
	if h.left.left.isRed() {
		h = t.rotateRight(h)
		t.flipColors(h)
	}
	return h
}

func (t *Tree) balance(h *Node) *Node {
	if h.right.isRed() {
		h = t.rotateLeft(h)
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = t.rotateRight(h)
	}
	if h.left.isRed() && h.right.isRed() {
		t.flipColors(h)
	}
	h.n = 1 + h.left.size() + h.right.size()
	return h
}

func (t *Tree) deleteMin(h *Node) *Node {
	if h.left == nil {
		return nil
	}
	if !h.left.isRed() && h.left.left.isRed() {
		h = t.moveRedLeft(h)
	}
	h.left = t.deleteMin(h.left)
	return t.balance(h)
}

func (t *Tree) deleteMax(h *Node) *Node {
	if h.left.isRed() {
		h = t.rotateRight(h)
	}
	if h.right == nil {
		return nil
	}

	if !h.right.isRed() && !h.right.left.isRed() {
		h = t.moveRedRight(h)
	}

	h.right = t.deleteMax(h.right)
	return t.balance(h)
}

// min returns the min node in tree.
func (t *Tree) min(x *Node) *Node {
	if x.left == nil {
		return x
	}
	return t.min(x.left)
}

func (t *Tree) max(x *Node) *Node {
	if x.right == nil {
		return x
	}
	return t.max(x.right)
}

func (t *Tree) delete(h *Node, key interface{}) *Node {
	if t.Comparator(key, h.key) < 0 {
		if !h.left.isRed() && !h.left.left.isRed() {
			h = t.moveRedLeft(h)
		}
		h.left = t.delete(h.left, key)
	} else {
		if h.left.isRed() {
			h = t.rotateRight(h)
		}

		if t.Comparator(key, h.key) == 0 && h.right == nil {
			return nil
		}

		if !h.right.isRed() && !h.right.left.isRed() {
			h = t.moveRedRight(h)
		}

		if t.Comparator(key, h.key) == 0 {
			x := t.min(h.right)
			h.key = x.key
			h.value = x.value
			h.right = t.deleteMin(h.right)
		} else {
			h.right = t.delete(h.right, key)
		}
	}
	return t.balance(h)
}

func (t *Tree) put(h *Node, key interface{}, val interface{}) *Node {
	if h == nil {
		return &Node{
			key:   key,
			value: val,
			left:  nil,
			right: nil,
			color: RED,
			n:     1,
		}
	}

	cmp := t.Comparator(key, h.key)
	if cmp < 0 {
		h.left = t.put(h.left, key, val)
	} else if cmp > 0 {
		h.right = t.put(h.right, key, val)
	} else {
		h.value = val
	}

	if h.right.isRed() && !h.left.isRed() {
		h = t.rotateLeft(h)
	}

	if h.left.isRed() && h.left.left.isRed() {
		h = t.rotateRight(h)
	}

	if h.left.isRed() && h.right.isRed() {
		t.flipColors(h)
	}

	h.n = 1 + h.left.size() + h.right.size()
	return h
}

func (t *Tree) get(h *Node, key interface{}) interface{} {
	if h == nil {
		return nil
	}
	cmp := t.Comparator(key, h.key)

	if cmp < 0 {
		return t.get(h.left, key)
	} else if cmp > 0 {
		return t.get(h.right, key)
	} else {
		return h.value
	}
}

func (t *Tree) selects(x *Node, k int) interface{} {
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

func (t *Tree) rank(x *Node, key interface{}) int {
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

func (t *Tree) keysByIndex(x *Node, keys *[]interface{}, lo interface{}, hi interface{}, cur *int) {
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

func (t *Tree) floor(x *Node, key interface{}) *Node {
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

func (t *Tree) ceiling(x *Node, key interface{}) *Node {
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
