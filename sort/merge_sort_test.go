package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
	result := MergeSort(arr)
	for _, item := range result {
		t.Log(item)
	}
}
