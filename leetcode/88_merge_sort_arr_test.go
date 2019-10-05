package leetcode

import "testing"

func TestMerge(t *testing.T) {
	// s := make([]int, 3, 6)
	// s[0] = 1
	// s[1] = 2
	// s[2] = 3
	// s = append(s, 0, 0, 0)
	// s2 := []int{2, 5, 6}
	// merge(s, 3, s2, 3)
	// t.Log(s)
	s3 := []int{2, 0}
	s4 := []int{1}
	merge(s3, 1, s4, 1)
	t.Log(s3)
}
