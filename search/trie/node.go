package trie

type node struct {
	count    int
	data     rune
	children []*node
}

func newNode(char rune) *node {
	return &node{
		count:    0,
		data:     char,
		children: make([]*node, 26),
	}
}

func (n *node) hasChild(char rune) bool {
	index := getIndex(char)
	if n.children[index] == nil {
		return false
	}
	return true
}

func (n *node) addChild(char rune) *node {
	index := getIndex(char)
	if n.children[index] == nil {
		n.children[index] = newNode(char)
	}
	n.children[index].count++
	return n.children[index]
}

func (n *node) getChild(char rune) *node {
	index := getIndex(char)
	return n.children[index]
}

func (n *node) suggestChildren() (suggest []rune) {
	for _, node := range n.children {
		if node != nil {
			suggest = append(suggest, node.data)
		}
	}
	return
}

func getIndex(char rune) rune {
	return char - 96
}
