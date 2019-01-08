package tree

const (
	RED   = true
	BLACK = false
)

type rbnode struct {
	key         Key
	val         Value
	left, right *rbnode
	color       bool
	n           int
}

func (x *rbnode) isRed() bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func (x *rbnode) size() int {
	if x == nil {
		return 0
	}
	return x.n
}

func (h *rbnode) rotateLeft() *rbnode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

func (h *rbnode) rotateRight() *rbnode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.n = h.n
	h.n = 1 + h.left.size() + h.right.size()
	return x
}

// flipColors converted the color of two red child nodes.
func (h *rbnode) flipColors() {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (x *rbnode) get(key Key) Value {
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

func (h *rbnode) put(key Key, value Value) *rbnode {
	if h == nil {
		return &rbnode{
			key:   key,
			val:   value,
			left:  nil,
			right: nil,
			color: RED,
			n:     1,
		}
	}

	cmp := key.CompareTo(h.key)
	if cmp < 0 {
		h.left = h.left.put(key, value)
	} else if cmp > 0 {
		h.right = h.right.put(key, value)
	} else {
		h.val = value
	}

	if h.right.isRed() && !h.left.isRed() {
		h = h.rotateLeft()
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = h.rotateRight()
	}
	if h.left.isRed() && h.right.isRed() {
		h.flipColors()
	}

	h.n = 1 + h.left.size() + h.right.size()
	return h
}

func (h *rbnode) moveRedLeft() *rbnode {
	h.flipColors()
	if h.right.left.isRed() {
		h.right = h.right.rotateRight()
		h = h.rotateLeft()
		h.flipColors()
	}
	return h
}
func (h *rbnode) moveRedRight() *rbnode {
	h.flipColors()
	if h.left.left.isRed() {
		h = h.rotateRight()
		h.flipColors()
	}
	return h
}

func (h *rbnode) balance() *rbnode {
	if h.right.isRed() {
		h = h.rotateLeft()
	}
	if h.left.isRed() && h.left.left.isRed() {
		h = h.rotateRight()
	}
	if h.left.isRed() && h.right.isRed() {
		h.flipColors()
	}

	h.n = 1 + h.left.size() + h.right.size()
	return h
}

func (h *rbnode) deleteMin() *rbnode {
	if h.left == nil {
		return nil
	}
	if !h.left.isRed() && h.left.left.isRed() {
		h = h.moveRedLeft()
	}
	h.left = h.left.deleteMin()
	return h.balance()
}

func (h *rbnode) deleteMax() *rbnode {
	if h.left.isRed() {
		h = h.rotateRight()
	}
	if h.right == nil {
		return nil
	}
	if !h.right.isRed() && !h.right.left.isRed() {
		h = h.moveRedRight()
	}

	h.right = h.right.deleteMax()
	return h.balance()
}

func (h *rbnode) min() *rbnode {
	if h.left == nil {
		return h
	}
	return h.left.min()
}

func (h *rbnode) delete(key Key) *rbnode {
	if key.CompareTo(h.key) < 0 {
		if !h.left.isRed() && !h.left.left.isRed() {
			h = h.moveRedLeft()
		}
		h.left = h.left.delete(key)
	} else {
		if h.left.isRed() {
			h = h.rotateRight()
		}

		if key.CompareTo(h.key) == 0 && (h.right == nil) {
			return nil
		}

		if !h.right.isRed() && !h.right.left.isRed() {
			h = h.moveRedRight()
		}

		if key.CompareTo(h.key) == 0 {
			x := h.right.min()
			h.key = x.key
			h.val = x.val
			h.right = h.right.deleteMin()
		} else {
			h.right = h.right.delete(key)
		}
	}
	return h.balance()
}

func (x *rbnode) keys(keys *[]Key, lo, hi Key, cur *int) {
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

type redBlackTree struct {
	root *rbnode
}

func NewRedBlackTree() *redBlackTree {
	return &redBlackTree{
		nil,
	}
}

func (t *redBlackTree) Put(key Key, val Value) {
	t.root = t.root.put(key, val)
	t.root.color = BLACK
}

func (t *redBlackTree) Get(key Key) Value {
	return t.root.get(key)
}

func (t *redBlackTree) IsEmpty() bool {
	return t.root == nil
}

func (t *redBlackTree) DeleteMin() {
	if t.IsEmpty() {
		return
	}

	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}
	t.root = t.root.deleteMin()
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func (t *redBlackTree) DeleteMax() {
	if t.IsEmpty() {
		return
	}

	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}
	t.root = t.root.deleteMax()
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func (t *redBlackTree) Delete(key Key) {
	if key == nil {
		return
	}

	if !t.Contain(key) {
		return
	}
	if !t.root.left.isRed() && !t.root.right.isRed() {
		t.root.color = RED
	}

	t.root = t.root.delete(key)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func (t *redBlackTree) Contain(key Key) bool {
	return t.Get(key) != nil
}
