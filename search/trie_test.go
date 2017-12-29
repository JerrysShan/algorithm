package search

import (
	"fmt"
	"testing"
)

var testRoot *TrieNode = createTrieNode()

func TestTrieInsert(t *testing.T) {
	t.Log("test TrieInsert")
	TrieInsert(testRoot, "apple")
	TrieInsert(testRoot, "application")
	t.Log("should app")
}

func TestTrieSearch(t *testing.T) {
	t.Log("test TrieSearch")
	ret := TrieSearch(testRoot, "applee")
	fmt.Println(ret, "------------->")
	if ret {
		t.Log("'app' find in trieRoot")
	} else {
		t.Log("'app' is not founded")
	}
	t.Log("ret should be true", ret)
}
