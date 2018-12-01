package sort

func QuickSort(arr []int, low, high int) {
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
		if i < j {
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[low], arr[i] = arr[i], arr[low]
	QuickSort(arr, low, i-1)
	QuickSort(arr, i+1, high)
}
