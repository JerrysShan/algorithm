package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	t.Log(majorityElement([]int{3, 2, 3}))
	t.Log(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
