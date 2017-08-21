package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 8, 10}
	index := binSearch(arr, 6)
	fmt.Println(index)
}

//数组必须是有序的
func binSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		if arr[low] == key {
			return low
		}
		if arr[high] == key {
			return high
		}
		mid := (high-low)/2 + low
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
