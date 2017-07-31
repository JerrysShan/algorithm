package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 5, 8, 10}
	index := binSearch(arr, 5)
	fmt.Println(index)
}

func binSearch(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	mid := (high - low) / 2
	for low <= high {
		if arr[low] == key {
			return low
		}
		if arr[high] == key {
			return high
		}
		if arr[mid] == key {
			break
		} else if arr[mid] > key {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (high-low)/2 + low
	}
	if low > high {
		return -1
	}
	return mid
}
