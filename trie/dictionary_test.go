package trie

import (
	"bufio"
	"io"
	"testing"

	"bytes"
	"errors"
	"os"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, len(arg), d.MaxLength)
}

func TestDictionaryShouldCallTrieAddAsTheNumberOfLoadedSetOfString(t *testing.T) {
	trie := new(MockObject)
	trie.On("Add", mock.AnythingOfType("string"))

	arg := []string{"test", "cat", "dog"}
	d := NewDict()
	d.trie = trie

	d.LoadStringSet(arg)

	trie.AssertNumberOfCalls(t, "Add", len(arg))
	assert.Equal(t, 4, d.MaxLength)
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
	assert.Equal(t, 0, d.MaxLength)
}

func TestDictionaryShouldReturnErrorIfLoadFileFail(t *testing.T) {
	OpenFile = func(f string) (*os.File, error) {
		return nil, errors.New("error")
	}

	d := NewDict()
	e := d.LoadFile("somefile")

	assert.Error(t, e)
}

func TestDictionaryShouldReturnFileContentAsStringArray(t *testing.T) {
	OpenFile = func(f string) (*os.File, error) {
		return os.NewFile(0, "test"), nil
	}

	GetScanner = func(f io.Reader) *bufio.Scanner {
		buf := bytes.NewBufferString("test\ndata")
		return bufio.NewScanner(buf)
	}

	d := NewDict()
	e := d.LoadFile("somefile")
	r1 := d.Has("data")
	r2 := d.Has("test")
	d.Clear()

	assert.Nil(t, e)
	assert.True(t, r1, "Text data should exist in dictionary, but not")
	assert.True(t, r2, "Text test should exist in dictionary, but not")
}
