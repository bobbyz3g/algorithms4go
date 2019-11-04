package trie

import (
	"github.com/Kaiser925/algorithms4go/queue"
)

// RuneTrie is of runes.
type RuneTrie struct {
	size  int
	value interface{}
	next  map[rune]*RuneTrie
}

// NewRuneTrie constructs and returns a new *RuneTrie
func NewRuneTrie() *RuneTrie {
	return &RuneTrie{
		next: make(map[rune]*RuneTrie),
	}
}

// Put inserts key-value into the tree, if there is an
// existing value, Put will replaces it.
func (t *RuneTrie) Put(key string, value interface{}) {
	node := t
	for _, k := range key {
		child := node.next[k]
		if child == nil {
			child = NewRuneTrie()
			node.next[k] = child
		}
		node = child
	}
	node.value = value
	t.size++
}

// Get returns value by its key or nil if key is not found in tree.
func (t *RuneTrie) Get(key string) interface{} {
	node := t.get(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *RuneTrie) get(key string) *RuneTrie {
	node := t
	for _, k := range key {
		node = node.next[k]
		if node == nil {
			return nil
		}
	}
	return node
}

type keyPath struct {
	r    rune
	node *RuneTrie
}

// Delete deletes the key and its value.
func (t *RuneTrie) Delete(key string) bool {
	node := t
	path := make([]keyPath, len(key))

	for i, k := range key {
		path[i] = keyPath{r: k, node: node}
		node = node.next[k]

		if node == nil {
			return false // the key does not exsit.
		}
	}

	node.value = nil

	if len(node.next) == 0 {
		for i := len(key) - 1; i >= 0; i-- {
			preNode := path[i].node
			r := path[i].r
			delete(preNode.next, r)
			if preNode.value != nil || len(preNode.next) > 0 {
				break
			}
		}
	}

	t.size--
	return true // success delete the key.
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *RuneTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

// Keys returns a Queue<string>
func (t *RuneTrie) Keys() *queue.Queue {
	return t.KeyWithPrefix("")
}

func getkeys(t *RuneTrie, pre string, keys *queue.Queue) {
	if t == nil {
		return
	}
	if t.value != nil {
		keys.Enqueue(pre)
	}

	for k := range t.next {
		getkeys(t.next[k], pre+string(k), keys)
	}
}

// IsEmpty returns the tree is empty or not.
func (t *RuneTrie) IsEmpty() bool {
	return t.size == 0
}

// LongestPrefixOf returns the longest key in the prefix of str.
func (t *RuneTrie) LongestPrefixOf(str string) string {
	return ""
}

// Size returns the number of key-value.
func (t *RuneTrie) Size() int {
	return t.size
}

// KeyWithPrefix returns all keys prefixed with str.
func (t *RuneTrie) KeyWithPrefix(str string) *queue.Queue {
	keys := queue.NewQueue()
	pre := str
	node := t.get(pre)

	getkeys(node, pre, keys)

	return keys
}

// KeyThatMatch returns the key matched with str.
// Note: "."can match any keyboard.
func (t *RuneTrie) KeyThatMatch(str string) *queue.Queue {
	return nil
}
