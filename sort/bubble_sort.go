package main

import "fmt"

func main() {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
	bubbleSort(arr)
	for _, item := range arr {
		fmt.Println(item)
	}
}

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j <length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
