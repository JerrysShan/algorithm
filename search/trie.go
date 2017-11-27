package search

type trieNode struct {
	data       string
	freq       int
	childNodes []trieNode
}
