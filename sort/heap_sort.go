package main

import "fmt"

func main() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heapSort(arr)
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}

func adjustHeap(arr []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left <= length && arr[left] > arr[i] {
		max = left
	}
	if right <= length && arr[right] > arr[max] {
		max = right
	}
	if max != i {
		arr[i], arr[max] = arr[max], arr[i]
		adjustHeap(arr, max, length)
	}
}

func buildHeap(arr []int) {
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, length-1)
	}
}

func heapSort(arr []int) {
	length := len(arr)
	buildHeap(arr)
	for i := length - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i-1)
	}
}
