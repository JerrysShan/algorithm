package sort

func adjustHeap(arr []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	min := i
	if left <= length && arr[left] < arr[i] {
		min = left
	}
	if right <= length && arr[right] < arr[min] {
		min = right
	}
	if min != i {
		arr[i], arr[min] = arr[min], arr[i]
		adjustHeap(arr, min, length)
	}
}

func buildHeap(arr []int) {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, length-1)
	}
}

func HeapSort(arr []int) {
	length := len(arr)
	buildHeap(arr)
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i-1)
	}
}
