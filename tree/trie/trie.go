package trie

type Trie struct {
}

// NewTrie allocates and returns a new Tire.
func NewTrie() *Trie {
	return &Trie{}
}

// Put inserts key-value into the tree, if there is an
// existing value, Put will replaces it.
func (t *Trie) Put(key string, value interface{}) {

}

// Get returns value by its key or nil if key is not found in tree.
func (t *Trie) Get(key string) interface{} {
	return nil
}

// Delete deletes the key and its value.
func (t *Trie) Delete(key string) {

}

// Contains returns true if tree contains key or false if doesn't contain.
func (t *Trie) Contains(key string) bool {
	return false
}

// Keys returns all keys.
func (t *Trie) Keys() []string {
	return nil
}

// IsEmpty returns the tree is empty or not.
func (t *Trie) IsEmpty() bool {
	return false
}

// LongestPrefixOf returns the longest key in the prefix of str.
func (t *Trie) LongestPrefixOf(str string) string {
	return ""
}

// Size returns the number of key-value.
func (t *Trie) Size() int {
	return 0
}

// KeyWithPrefix returns all keys prefixed with str.
func (t *Trie) KeyWithPrefix(str string) []string {
	return nil
}

// KeyThatMatch returns the key matched with str.
// Note: "."can match any keyboard.
func (t *Trie) KeyThatMatch(str string) []string {
	return nil
}
