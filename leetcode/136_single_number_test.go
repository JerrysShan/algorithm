package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	s1 := []int{2, 2, 1}
	t.Log(singleNumber(s1))
	s2 := []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(s2))
}
