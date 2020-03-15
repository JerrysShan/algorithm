package offer

import "testing"

func TestMaxValue(t *testing.T) {
	arr := make([][]int, 2)
	arr[0] = []int{1, 2, 5}
	arr[1] = []int{3, 2, 1}
	// arr[2] = []int{4, 2, 1}
	t.Log(maxValue(arr))
}
