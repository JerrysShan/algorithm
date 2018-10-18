package search

import "testing"

func TestLongestPrefix(t *testing.T) {
	strs := []string{"flight", "flow", "flower"}
	prefix := LongestPrefix(strs)
	t.Log(prefix)
}
