package leetcode

import "testing"

func TestCanThreePartsEqualsSum(t *testing.T) {
	t.Log(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	t.Log(canThreePartsEqualSum([]int{10, -10, 10, -10, 10, -10, 10, -10}))
}
