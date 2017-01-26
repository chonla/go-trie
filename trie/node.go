package trie

// Node is trie node
type Node struct {
	Children childNodes
	Value    string
	Terminal bool
}

type childNodes map[string]*Node

// NewNode is to create a new trie node
func NewNode(v string) *Node {
	return &Node{Value: v, Children: childNodes{}, Terminal: false}
}

// HasChild returns true if node has child with given index
func (n *Node) HasChild(v string) bool {
	if n.Children == nil {
		return false
	}

	if _, ok := n.Children[v]; ok {
		return true
	}
	return false
}

// Clear cleans up children
func (n *Node) Clear() {
	for k := range n.Children {
		n.Children[k].Clear()

		delete(n.Children, k)
	}
}
