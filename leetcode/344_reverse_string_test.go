package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	t.Logf("%s", s)
}
