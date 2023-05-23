package trie

type Node struct {
	end      bool
	value    string
	children []*Node
}

type Trie struct {
	node *Node
}

func New() *Trie {
	return &Trie{node: &Node{value: "", children: make([]*Node, 0)}}
}

func (t *Trie) Insert(word string) {
	insertAtNode(t.node, word)
}

func (t *Trie) Search(word string) bool {
	return retrieve(t.node, word, false)
}

func (t *Trie) StartsWith(prefix string) bool {
	return retrieve(t.node, prefix, true)
}

func retrieve(n *Node, word string, p bool) bool {
	if len(word) == 0 {
		if p {
			return true
		} else {
			return n.end
		}
	}

	l := word[:1]
	for _, v := range n.children {
		if v.value == l {
			return retrieve(v, word[1:], p)
		}
	}

	return false
}

func insertAtNode(n *Node, word string) {
	if len(word) == 0 {
		n.end = true
		return
	}

	l := word[:1]

	for _, v := range n.children {
		if v.value == l {
			insertAtNode(v, word[1:])
			return
		}
	}

	nn := Node{value: l, children: make([]*Node, 0)}
	n.children = append(n.children, &nn)
	insertAtNode(&nn, word[1:])
}
