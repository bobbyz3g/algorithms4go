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

func TestPutAndGet(t *testing.T) {
	trie := NewRuneTrie()

	for k, v := range testInput {
		trie.Put(k, v)
	}

	if al, l := trie.Size(), len(testInput); al != l {
		t.Errorf("Trie get error: excepted: %d, got %d", l, al)
	}

	for k, v := range testInput {
		assert.Equal(t, v, trie.Get(k))
	}

	assert.Equal(t, true, trie.Contains("Sells"))
}
