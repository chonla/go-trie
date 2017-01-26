package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNodeShouldReturnEmptyNode(t *testing.T) {
	n := NewNode("")

	assert.Equal(t, "", n.Value, "Default value of trie node should be empty string.")
	assert.False(t, n.Terminal, "Default value of terminal should be false.")
	assert.Equal(t, 0, len(n.Children), "Default number of children should be zero.")
}

func TestHasChildShouldReturnFalseWithEmptyNode(t *testing.T) {
	n := NewNode("")

	assert.False(t, n.HasChild("a"), "HasChild should return false when child does not exist with given branch key, but true.")
}

func TestDirectInstanceNodeHasChildShouldReturnFalseWithEmptyNode(t *testing.T) {
	n := &Node{}

	assert.False(t, n.HasChild("a"), "HasChild should return false when child does not exist with given branch key, but true.")
}

func TestHasChildShouldReturnTrueIfGivenChildExist(t *testing.T) {
	n := NewNode("")
	n.Children["a"] = NewNode("")

	assert.True(t, n.HasChild("a"), "HasChild should return true when child exist with given branch key, but not.")

	n.Clear()
}

func TestHasChildShouldReturnFalseIfGivenChildDoesNotExist(t *testing.T) {
	n := NewNode("")
	n.Children["a"] = NewNode("")

	assert.False(t, n.HasChild("b"), "HasChild should return false when child does not exist with given branch key, but true.")

	n.Clear()
}

func TestClearShouldClearAllChildren(t *testing.T) {
	n := NewNode("")
	n.Children["a"] = NewNode("")
	n.Clear()

	assert.Equal(t, 0, len(n.Children), "Number of children after node is clear should be zero.")
}
