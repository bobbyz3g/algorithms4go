package trie

import (
	"github.com/Kaiser925/algorithms4go/queue"
	"reflect"
)

// RuneTrie is of runes.
type RuneTrie[T any] struct {
	size  int
	value T
	next  map[rune]*RuneTrie[T]
}

// NewRuneTrie constructs and returns a new *RuneTrie
func NewRuneTrie[T any]() *RuneTrie[T] {
	return &RuneTrie[T]{
		next: make(map[rune]*RuneTrie[T]),
	}
}

// Put inserts key-value into the tree, if there is an
// existing value, Put will replaces it.
func (t *RuneTrie[T]) Put(key string, value T) {
	node := t
	for _, k := range key {
		child := node.next[k]
		if child == nil {
			child = NewRuneTrie[T]()
			node.next[k] = child
		}
		node = child
	}
	node.value = value
	t.size++
}

// Get returns value by its key or nil if key is not found in tree.
func (t *RuneTrie[T]) Get(key string) (T, bool) {
	var noop T
	if key == "" {
		return noop, false
	}
	node := t.get(key)
	if node == nil {
		return noop, false
	}
	return node.value, true
}

func (t *RuneTrie[T]) get(key string) *RuneTrie[T] {
	node := t
	for _, k := range key {
		node = node.next[k]
		if node == nil {
			return nil
		}
	}
	return node
}

type keyPath[T any] struct {
	r    rune
	node *RuneTrie[T]
}

// Delete deletes the key and its value.
func (t *RuneTrie[T]) Delete(key string) bool {
	node := t
	path := make([]keyPath[T], len(key))

	for i, k := range key {
		path[i] = keyPath[T]{r: k, node: node}
		node = node.next[k]

		if node == nil {
			return false // the key does not exsit.
		}
	}
	var noop T
	node.value = noop

	if len(node.next) == 0 {
		for i := len(key) - 1; i >= 0; i-- {
			preNode := path[i].node
			r := path[i].r
			delete(preNode.next, r)
			if reflect.DeepEqual(preNode.value, noop) || len(preNode.next) > 0 {
				break
			}
		}
	}

	t.size--
	return true // success delete the key.
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *RuneTrie[T]) Contains(key string) bool {
	_, ok := t.Get(key)
	return ok
}

// Keys returns a Queue<string>
func (t *RuneTrie[T]) Keys() *queue.Queue[string] {
	return t.KeyWithPrefix("")
}

func getkeys[T any](t *RuneTrie[T], pre string, keys *queue.Queue[string]) {
	if t == nil {
		return
	}
	var noop T
	if !reflect.DeepEqual(t.value, noop) {
		keys.Enqueue(pre)
	}

	for k := range t.next {
		getkeys(t.next[k], pre+string(k), keys)
	}
}

// IsEmpty returns the tree is empty or not.
func (t *RuneTrie[T]) IsEmpty() bool {
	return t.size == 0
}

// Size returns the number of key-value.
func (t *RuneTrie[T]) Size() int {
	return t.size
}

// KeyWithPrefix returns all keys prefixed with str.
func (t *RuneTrie[T]) KeyWithPrefix(str string) *queue.Queue[string] {
	keys := queue.NewQueue[string]()
	pre := str
	node := t.get(pre)

	getkeys(node, pre, keys)

	return keys
}
