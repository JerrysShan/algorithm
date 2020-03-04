package leetcode

import "testing"

func TestFindDuplicate(t *testing.T) {
	t.Log(findDuplicate([]int{1, 3, 4, 2, 2}))
	t.Log(findDuplicate([]int{3, 1, 3, 4, 2}))
}
