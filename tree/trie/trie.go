package trie

// RuneTrie is of runes.
type RuneTrie struct {
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
}

// Get returns value by its key or nil if key is not found in tree.
func (t *RuneTrie) Get(key string) interface{} {
	node := t
	for _, k := range key {
		node = node.next[k]
		if node == nil {
			return nil
		}
	}
	return node.value
}

// Delete deletes the key and its value.
func (t *RuneTrie) Delete(key string) {
}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *RuneTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

// Keys returns all keys.
func (t *RuneTrie) Keys() []string {
	return nil
}

// IsEmpty returns the tree is empty or not.
func (t *RuneTrie) IsEmpty() bool {
	return len(t.next) == 0 && t.value == nil
}

// LongestPrefixOf returns the longest key in the prefix of str.
func (t *RuneTrie) LongestPrefixOf(str string) string {
	return ""
}

// Size returns the number of key-value.
func (t *RuneTrie) Size() int {
	node := t
	cnt := 0

	if node.value != nil {
		cnt++
	}

	for k := range node.next {
		cnt += node.next[k].Size()
	}
	return cnt
}

// KeyWithPrefix returns all keys prefixed with str.
func (t *RuneTrie) KeyWithPrefix(str string) []string {
	return nil
}

// KeyThatMatch returns the key matched with str.
// Note: "."can match any keyboard.
func (t *RuneTrie) KeyThatMatch(str string) []string {
	return nil
}
