package leetcode

import "testing"

func TestMaxSubArrary(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Log(maxSubArray(arr))
}
