package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
	QuickSort(arr, 0, len(arr)-1)
	for _, item := range arr {
		t.Log(item)
	}
}
