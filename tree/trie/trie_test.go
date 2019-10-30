package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = map[string]string{
	"Bob":    "bob",
	"Alice":  "Alice",
	"Sea":    "sea",
	"Test":   "test",
	"Sells":  "sells",
	"Shells": "Shells",
}

func constructTestTrie(testInput map[string]string) *RuneTrie {
	trie := NewRuneTrie()

	for k, v := range testInput {
		trie.Put(k, v)
	}

	return trie
}

func TestPutAndGet(t *testing.T) {
	trie := constructTestTrie(testInput)

	if al, l := trie.Size(), len(testInput); al != l {
		t.Errorf("Trie get error: excepted: %d, got %d", l, al)
	}

	for k, v := range testInput {
		assert.Equal(t, v, trie.Get(k))
	}

	assert.Equal(t, true, trie.Contains("Sells"))
}

func TestDelete(t *testing.T) {
	trie := constructTestTrie(testInput)

	assert.Equal(t, true, trie.Delete("Sells"))
	assert.Equal(t, false, trie.Delete("Sells"))
	assert.Equal(t, len(testInput)-1, trie.Size())
	assert.Equal(t, nil, trie.Get("Sells"))
	assert.Equal(t, testInput["Sea"], trie.Get("Sea"))
}
