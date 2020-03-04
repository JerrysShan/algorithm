package leetcode

import "testing"

func TestTopKFrequent(t *testing.T) {
	t.Log(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	t.Log(topKFrequent([]int{1}, 1))
}
