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
	tt := NewRuneTrie()

	assert.Equal(t, 0, tt.Size())
	trie := constructTestTrie(testInput)

	if al, l := trie.Size(), len(testInput); al != l {
		t.Errorf("Trie get error: excepted: %d, got %d", l, al)
	}

	for k, v := range testInput {
		assert.Equal(t, v, trie.Get(k))
	}

	assert.Equal(t, true, trie.Contains("Sells"))
	assert.Equal(t, nil, trie.Get(""))
}

func TestDelete(t *testing.T) {
	trie := constructTestTrie(testInput)

	assert.Equal(t, true, trie.Delete("Sells"))
	assert.Equal(t, false, trie.Delete("Sells"))
	assert.Equal(t, len(testInput)-1, trie.Size())
	assert.Equal(t, nil, trie.Get("Sells"))
	assert.Equal(t, testInput["Sea"], trie.Get("Sea"))
}

func TestKeys(t *testing.T) {
	trie := NewRuneTrie()

	assert.Equal(t, true, trie.Keys().Empty())
	trie.Put("a", 1)

	keys := trie.Keys()
	assert.Equal(t, 1, keys.Len())
	k, _ := keys.Dequeue()
	assert.Equal(t, "a", k)

	trie = constructTestTrie(testInput)

	keys = trie.Keys()

	iter := keys.Iter()

	for iter.HasNext() {
		key, _ := iter.Value()
		_, has := testInput[key.(string)]
		assert.Equal(t, true, has)
	}
}

func TestKeyWithPrefix(t *testing.T) {
	trie := constructTestTrie(testInput)

	keys := trie.KeyWithPrefix("Se")
	assert.Equal(t, 2, keys.Len())

	for iter := keys.Iter(); iter.HasNext(); {
		if key, _ := iter.Value(); key.(string) == "Sea" || key.(string) == "Sells" {

		} else {
			t.Error("KeyWithPrefix error")
		}
	}

	keys = trie.KeyWithPrefix("Sellls")
	assert.Equal(t, 0, keys.Len())
}
