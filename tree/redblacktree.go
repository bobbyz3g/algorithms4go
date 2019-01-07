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

// flipColor converted the color of two red child nodes.
// 1 4-node => 3 2-node.
func (h *rbnode) flipColor() {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
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
		h.flipColor()
	}

	h.n = 1 + h.left.size() + h.right.size()
	return h
}

type redBlackTree struct {
	root *rbnode
}

func NewRedBlackTree(key Key, val Value) *redBlackTree {
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
