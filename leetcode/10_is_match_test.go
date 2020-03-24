package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	t.Log(isMatch("aa", "a"))
	t.Log(isMatch("ab", ".*"))
}
