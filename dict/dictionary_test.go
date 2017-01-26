package dict

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Add(v string) {
	m.Called(v)
}

func (m *MockObject) Has(v string) bool {
	m.Called(v)
	return true
}

func (m *MockObject) Clear() {
	m.Called()
}

func TestDictionaryShouldCallTrieAddWhenLoadNewString(t *testing.T) {
	trie := new(MockObject)
	trie.On("Add", mock.AnythingOfType("string"))

	arg := "test"
	d := NewDict()
	d.trie = trie

	d.LoadString(arg)

	trie.AssertCalled(t, "Add", arg)
}

func TestDictionaryShouldCallTrieAddAsTheNumberOfLoadedSetOfString(t *testing.T) {
	trie := new(MockObject)
	trie.On("Add", mock.AnythingOfType("string"))

	arg := []string{"test", "cat", "dog"}
	d := NewDict()
	d.trie = trie

	d.LoadStringSet(arg)

	trie.AssertNumberOfCalls(t, "Add", len(arg))
}

func TestDictionaryShouldReturnTrieHas(t *testing.T) {
	trie := new(MockObject)
	trie.On("Has", mock.AnythingOfType("string"))

	arg := "test"
	d := NewDict()
	d.trie = trie

	d.Has(arg)

	trie.AssertCalled(t, "Has", arg)
}

func TestDictionaryShouldCallTrieClear(t *testing.T) {
	trie := new(MockObject)
	trie.On("Clear")

	d := NewDict()
	d.trie = trie

	d.Clear()

	trie.AssertCalled(t, "Clear")
}
