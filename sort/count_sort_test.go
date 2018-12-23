package sort

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := []int{9, 3, 5, 4, 9, 1, 2, 7, 8, 1, 3, 6, 5, 3, 4, 12, 10, 9, 7, 9}
	SimpleCountSort(arr)
	t.Log(arr)
}
