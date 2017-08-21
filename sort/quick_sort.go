package main

import "fmt"

var count int

func main() {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
	quickSort(arr, 0, len(arr)-1)
	for _, item := range arr {
		fmt.Println(item)
	}
}

func quickSort(arr []int, low, high int) {
	if low > high {
		return
	}
	key := arr[low]
	i := low
	j := high
	for i < j {
		for arr[j] >= key && i < j {
			j--
		}
		for arr[i] <= key && i < j {
			i++
		}
		fmt.Print("j is", " ", j, " ", count, " | ")
		fmt.Print("i is", " ", i, " ", count, " | ")
		count++
		if i < j {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[low], arr[i] = arr[i], arr[low]
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}
