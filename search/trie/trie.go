package trie

import (
	"fmt"
	"strings"
)

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode('*'),
	}
}

func (t *Trie) AddWord(word string) {
	word = strings.ToLower(word)
	cur := t.root
	for _, char := range word {
		cur = cur.addChild(char)
	}
}

func (t *Trie) ExistWord(word string) bool {
	word = strings.ToLower(word)
	node := t.getLastCharacterChild(word)
	if node == nil {
		return false
	}
	return true
}

func (t *Trie) SuggestWords(word string) (suggest []string) {
	word = strings.ToLower(word)
	cur := t.getLastCharacterChild(word)
	if cur == nil {
		return
	}
	temp := cur.suggestChildren()
	for _, char := range temp {
		suggest = append(suggest, fmt.Sprintf("%c", char))
	}
	return
}

func (t *Trie) getLastCharacterChild(word string) *node {
	cur := t.root
	for _, char := range word {
		if !cur.hasChild(char) {
			return nil
		}
		cur = cur.getChild(char)
	}
	return cur
}

func Show(n *node) {
	if n == nil {
		return
	}
	for _, n := range n.children {
		if n != nil {
			fmt.Printf("%c:%d \n", n.data, n.count)
			Show(n)
		}
	}
}
