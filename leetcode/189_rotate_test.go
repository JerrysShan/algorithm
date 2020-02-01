package leetcode

import "testing"

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	t.Log(arr)
	arr2 := []int{-1, -100, 3, 99}
	rotate(arr2, 2)
	t.Log(arr2)
}
