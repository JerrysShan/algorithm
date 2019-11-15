package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	t.Log(maxProduct([]int{2, 3, -2, 4}))
	t.Log(maxProduct([]int{-2, 3, -4}))
	t.Log(maxProduct([]int{0, 2}))
}
