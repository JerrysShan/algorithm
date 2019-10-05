package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	t.Log(isValid("()[]{}"))
	t.Log(isValid("(]"))
	t.Log(isValid("]("))
	t.Log(isValid("{[]}"))
}
