package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("flower")
	trie.AddWord("flow")
	trie.AddWord("flight")
	Show(trie.root)
	t.Log(trie.ExistWord("flower"))
	t.Log(trie.ExistWord("dog"))
	t.Log(trie.SuggestWords("fl"))
	t.Log(trie.SuggestWords("fi"))
}
