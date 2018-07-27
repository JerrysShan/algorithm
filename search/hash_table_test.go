package search

import (
	"fmt"
	"testing"
)

func TestHashTable(t *testing.T) {
	t.Log("test hashtable")
	h := NewHashTable()
	h.AddKV("add", "ssk")
	h.AddKV("bdc", "jerry")
	h.AddKV("123", "jerry")
	fmt.Println(h.GetVal("add"))
	fmt.Println(h.GetVal("bdc"))
	fmt.Println(h.GetVal("123"))
}
