package tree

type node struct {
	key         Key
	val         Value
	left, right *node // Pointer to subtree
	n           int   //The total number of nodes in the subtree rooted at this node.
}

func (x *node) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

func (x *node) min() *node {
	if x.left == nil {
		return x
	}
	return x.left.min()
}

func (x *node) max() *node {
	if x.right == nil {
		return x
	}
	return x.right.max()
}

func (x *node) get(key Key) Value {
	if x == nil {
		return nil
	}

	cmp := key.CompareTo(x.key)

	if cmp < 0 {
		return x.left.get(key)
	} else if cmp > 0 {
		return x.right.get(key)
	} else {
		return x.val
	}
}

func (x *node) put(key Key, val Value) *node {
	if x == nil {
		return newNode(key, val, 1)
	}

	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = x.left.put(key, val)
	} else if cmp > 0 {
		x.right = x.right.put(key, val)
	} else {
		x.val = val
	}
	x.n = x.left.size() + x.right.size() + 1
	return x
}

func (x *node) floor(key Key) *node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)

	if cmp == 0 {
		return x
	} else if cmp < 0 {
		return x.left.floor(key)
	}

	t := x.right.floor(key)
	if t != nil {
		return t
	} else {
		return x
	}
}

func (x *node) keySelect(k int) Key {
	if x == nil {
		return nil
	}

	t := x.left.size()
	if t > k {
		return x.left.keySelect(k)
	} else if t < k {
		return x.right.keySelect(k)
	} else {
		return x.key
	}
}

func (x *node) rank(key Key) int {
	if x == nil {
		return 0
	}
	cmp := key.CompareTo(x.key)

	if cmp < 0 {
		return x.left.rank(key)
	} else if cmp > 0 {
		return 1 + x.left.size() + x.right.rank(key)
	} else {
		return x.left.size()
	}
}

func (x *node) deleteMin() *node {
	if x.left == nil {
		return x.right
	}
	x.left = x.left.deleteMin()
	x.n = x.left.size() + x.right.size() + 1
	return x
}

func (x *node) delete(key Key) *node {
	if x == nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = x.left.delete(key)
	} else if cmp > 0 {
		x.right = x.right.delete(key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}

		t := x
		x = t.right.min()
		x.right = x.right.deleteMin()
		x.left = t.left
	}
	x.n = x.left.size() + x.right.size() + 1
	return x
}

func (x *node) keys(keys *[]Key, lo, hi Key, cur *int) {
	if x == nil {
		return
	}
	cmplo := lo.CompareTo(x.key)
	cmphi := hi.CompareTo(x.key)
	if cmplo < 0 {
		x.left.keys(keys, lo, hi, cur)
	}

	if cmplo <= 0 && cmphi >= 0 {
		(*keys)[*cur] = x.key
		*cur++
	}

	if cmphi > 0 {
		x.right.keys(keys, lo, hi, cur)
	}
}

func newNode(key Key, val Value, n int) *node {
	return &node{
		key,
		val,
		nil,
		nil,
		n,
	}
}

type bst struct {
	root *node
}

func NewBST() *bst {
	return &bst{
		nil,
	}
}

// Size return the number of nodes in bst.
func (t *bst) Size() int {
	return t.root.size()
}

// Get returns the value of the "key" in bst, if there is no "key", will return nil.
func (t *bst) Get(key Key) Value {
	return t.root.get(key)
}

// Put put key-value to bst, if the key is already in bst,
// Put wil update the value of the key.
func (t *bst) Put(key Key, val Value) {
	t.root = t.root.put(key, val)
}

func (t *bst) Min() Key {
	if t.root == nil {
		return nil
	}
	return t.root.min().key
}

func (t *bst) Max() Key {
	if t.root == nil {
		return nil
	}
	return t.root.max().key
}

func (t *bst) Floor(key Key) Key {
	x := t.root.floor(key)
	if x == nil {
		return nil
	}
	return x.key
}

func (t *bst) Select(k int) Key {
	return t.root.keySelect(k)
}

func (t *bst) Rank(key Key) int {
	return t.root.rank(key)
}

func (t *bst) Delete(key Key) {
	t.root = t.root.delete(key)
}

func (t *bst) KeysByIndex(lo, hi Key) []Key {
	keys := make([]Key, t.Size())
	cur := 0
	t.root.keys(&keys, lo, hi, &cur)
	return keys
}

func (t *bst) Keys() []Key {
	return t.KeysByIndex(t.Min(), t.Max())
}

type IntCmp int

func (i IntCmp) CompareTo(p interface{}) int {
	return (int)(i - p.(IntCmp))
}

type IntBST struct {
	baseTree *bst
}

func NewIntBST() *IntBST                       { return &IntBST{NewBST()} }
func (t *IntBST) Get(key int) Value            { return t.baseTree.Get(IntCmp(key)) }
func (t *IntBST) Put(key int, val interface{}) { t.baseTree.Put(IntCmp(key), val) }
func (t *IntBST) Min() int                     { return (int)(t.baseTree.Min().(IntCmp)) }
func (t *IntBST) Max() int                     { return (int)(t.baseTree.Max().(IntCmp)) }
func (t *IntBST) Size() int                    { return t.baseTree.Size() }
func (t *IntBST) Delete(key int)               { t.baseTree.Delete(IntCmp(key)) }
func (t *IntBST) Keys() []int                  { return t.KeysByIndex(t.baseTree.Min(), t.baseTree.Max()) }
func (t *IntBST) KeysByIndex(lo, hi Key) []int {
	a := t.baseTree.KeysByIndex(lo, hi)
	b := make([]int, len(a))
	for index, val := range a {
		b[index] = (int)(val.(IntCmp))
	}
	return b
}
