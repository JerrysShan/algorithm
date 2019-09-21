package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t.Log(removeDuplicates(arr))
}
