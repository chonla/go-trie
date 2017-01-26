package trie

// Trie is Trie structure
type Trie struct {
	Root *Node
}

// NewTrie is to create a new trie
func NewTrie() *Trie {
	return &Trie{Root: NewNode("")}
}

// Add is to import a string to trie
func (t *Trie) Add(v string) {
	ptr := t.Root

	t.add(ptr, v)
}

// Has return true if trie contains given string
func (t *Trie) Has(v string) bool {
	ptr := t.Root

	if len(v) > 0 {
		return t.has(ptr, v)
	}
	return false
}

func (t *Trie) add(n *Node, v string) {
	if len(v) > 0 {
		c := string([]rune(v)[0:1])
		if !n.HasChild(c) {
			t := n.Value + c
			m := NewNode(t)
			n.Children[c] = m
		}
		v = string([]rune(v)[1:])
		if len(v) > 0 {
			t.add(n.Children[c], v)
		} else {
			n.Children[c].Terminal = true
		}
	}
}

func (t *Trie) has(n *Node, v string) bool {
	c := string([]rune(v)[0:1])
	if !n.HasChild(c) {
		return false
	}
	v = string([]rune(v)[1:])
	if len(v) == 0 {
		return n.Children[c].Terminal
	}
	return t.has(n.Children[c], v)
}

// Clear to clear children nodes
func (t *Trie) Clear() {
	t.Root.Clear()
}
