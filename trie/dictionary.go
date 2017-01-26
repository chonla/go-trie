package trie

// Dictionary represent a dictionary
type Dictionary struct {
	trie ITrie
}

// NewDict is to create a new dictionary
func NewDict() *Dictionary {
	return &Dictionary{trie: NewTrie()}
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
}

// Has if a string is contained in trie
func (d *Dictionary) Has(t string) bool {
	return d.trie.Has(t)
}

// Clear is to clear dictionary
func (d *Dictionary) Clear() {
	d.trie.Clear()
}
