package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	t.Log(missingNumber([]int{3, 0, 1}))
	t.Log(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	t.Log(missingNumber([]int{0, 1}))
	t.Log(missingNumber([]int{1, 0}))
	t.Log(missingNumber([]int{0}))
	t.Log(missingNumber([]int{1}))

}
