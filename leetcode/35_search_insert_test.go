package leetcode

import "testing"

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 3, 5, 6}
	t.Log(searchInsert(arr, 5))
	t.Log(searchInsert(arr, 2))
	t.Log(searchInsert(arr, 7))
	t.Log(searchInsert(arr, 0))
}
