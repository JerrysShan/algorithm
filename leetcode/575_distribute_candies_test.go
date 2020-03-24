package leetcode

import "testing"

func TestDistributeCandies(t *testing.T) {
	t.Log(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	t.Log(distributeCandies([]int{1, 1, 2, 3}))
}
