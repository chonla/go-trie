package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTrieShouldReturnEmptyTrieWithARootInstance(t *testing.T) {
	trie := NewTrie()

	assert.NotNil(t, trie.Root, "Default root pointer should not be nil.")
}

func TestTrieAddTwoString(t *testing.T) {
	trie := NewTrie()

	trie.Add("good")
	trie.Add("cat")

	assert.True(t, trie.Root.HasChild("g"), "Root branch should contain child g.")
	assert.True(t, trie.Root.HasChild("c"), "Root branch should contain child c.")
	assert.False(t, trie.Root.HasChild("o"), "Root branch should not contain child o.")

	trie.Clear()
}

func TestHasShouldReturnTrueIfStringExists(t *testing.T) {
	trie := NewTrie()

	trie.Add("good")
	trie.Add("cat")
	trie.Add("goods")

	assert.True(t, trie.Has("good"), "Has should return true for existing string.")
	assert.True(t, trie.Has("cat"), "Has should return true for existing string.")
	assert.True(t, trie.Has("goods"), "Has should return true for existing string.")

	trie.Clear()
}

func TestHasShouldReturnFalseIfFindingEmptyString(t *testing.T) {
	trie := NewTrie()

	assert.False(t, trie.Has(""), "Has should return false for empty string.")

	trie.Clear()
}

func TestHasShouldReturnFalseIfStringDoesNotExists(t *testing.T) {
	trie := NewTrie()

	trie.Add("good")
	trie.Add("cat")
	trie.Add("goods")

	assert.False(t, trie.Has("bad"), "Has should return false for non existing string.")
	assert.False(t, trie.Has("goo"), "Has should return false for non existing substring.")

	trie.Clear()
}
