package search

type TrieNode struct {
	data       string
	freq       int
	childNodes []*TrieNode
}

func createTrieNode() *TrieNode {
	node := &TrieNode{
		freq:       0,
		childNodes: make([]*TrieNode, 26),
	}
	return node
}

func TrieInsert(root *TrieNode, key string) {
	if len(key) == 0 {
		return
	}
	chars := []byte(key)
	index := chars[0] - 'a'
	if root.childNodes[index] == nil {
		root.childNodes[index] = createTrieNode()
	}
	root.childNodes[index].data = key[:1]
	if len(key) == 1 {
		root.childNodes[index].freq++
	} else {
		TrieInsert(root.childNodes[index], key[1:])
	}
}

func TrieSearch(root *TrieNode, key string) bool {
	chars := []byte(key)
	p := root
	for _, c := range chars {
		index := c - 'a'
		if p.childNodes[index] == nil {
			return false
		}
		p = p.childNodes[index]
	}
	return true
}

func TrieShow(root *TrieNode, str string) {
	for _, v := range root.childNodes {
		if v != nil {
			str = str + "->" + v.data
			TrieShow(v, str)
		}
	}
}
