package leetcode

import "testing"

func TestSearchMatrix(t *testing.T) {
	arr := make([][]int, 7)

	arr[0] = []int{3, 3, 8, 13, 13, 18}
	arr[1] = []int{4, 5, 11, 13, 18, 20}
	arr[2] = []int{9, 9, 14, 15, 23, 23}
	arr[3] = []int{13, 18, 22, 22, 25, 27}
	arr[4] = []int{18, 22, 23, 28, 30, 33}
	arr[5] = []int{21, 25, 28, 30, 35, 35}
	arr[6] = []int{24, 25, 33, 36, 37, 40}

	// t.Log(searchMatrix(arr, 20))
	// t.Log(searchMatrix(arr, 5))
	t.Log(searchMatrix(arr, 21))
	// t.Log(searchMatrix(arr, 3))
}
