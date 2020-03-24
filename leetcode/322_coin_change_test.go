package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	// t.Log(coinChange([]int{1, 2, 5}, 11))
	// t.Log(coinChange([]int{2}, 3))

	t.Log(split([]int{1, 2, 5}, 10))
	t.Log(split([]int{2}, 3))
}
