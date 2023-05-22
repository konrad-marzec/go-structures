package trie

type Node struct {
	value    string
	children *[]Node
}

type Trie struct {
	node *Node
}

func New() *Trie {
	children := make([]Node, 0)

	return &Trie{node: &Node{value: "", children: &children}}
}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Search(word string) bool {
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return false
}
