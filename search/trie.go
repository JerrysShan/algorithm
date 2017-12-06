package search

type trieNode struct {
	data       string
	freq       int
	childNodes []*trieNode
}

func createTrieNode() *trieNode {
	node := &trieNode{
		freq:       0,
		childNodes: make([]*trieNode, 26),
	}
	return node
}

func trieInsert(root *trieNode, key string) {

}

func trieSearch(root *trieNode, key string) {

}
