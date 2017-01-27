package trie

import (
	"bufio"
	"os"
)

// OpenFile is os.Open wrapper, to ease mocking
var OpenFile = os.Open

// GetScanner is bufio.NewScanner wrapper, to ease mocking
var GetScanner = bufio.NewScanner

// ITrie defines interface for trie
type ITrie interface {
	Add(v string)
	Has(v string) bool
	Clear()
}

// Dictionary represent a dictionary
type Dictionary struct {
	trie  ITrie
	depth int
}

// NewDict is to create a new dictionary
func NewDict() *Dictionary {
	return &Dictionary{trie: NewTrie(), depth: 0}
}

// LoadStringSet is to load array of string into dictionary
func (d *Dictionary) LoadStringSet(ta []string) {
	for t := range ta {
		d.LoadString(ta[t])
	}
}

// LoadString is to load a string into dictionary
func (d *Dictionary) LoadString(t string) {
	d.trie.Add(t)
	l := len([]rune(t))
	if l > d.depth {
		d.depth = l
	}
}

// LoadFile is to load from a file
func (d *Dictionary) LoadFile(f string) error {
	l, e := d.readLines(f)

	if e != nil {
		return e
	}

	d.LoadStringSet(l)

	return nil
}

func (d *Dictionary) readLines(f string) ([]string, error) {
	file, err := OpenFile(f)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := GetScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// Has if a string is contained in trie
func (d *Dictionary) Has(t string) bool {
	return d.trie.Has(t)
}

// Clear is to clear dictionary
func (d *Dictionary) Clear() {
	d.trie.Clear()
	d.depth = 0
}

// Depth return max length of trie
func (d *Dictionary) Depth() int {
	return d.depth
}
