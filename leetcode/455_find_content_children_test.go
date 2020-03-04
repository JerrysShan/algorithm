package leetcode

import "testing"

func TestFindContentChildren(t *testing.T) {
	t.Log(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	t.Log(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
