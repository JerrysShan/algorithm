package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	HeapSort(arr)
	for _, val := range arr {
		t.Logf("%v ", val)
	}
}
