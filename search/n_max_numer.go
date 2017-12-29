package search

func FindNMax(arr []int, low, high, n int) int {

	q := 0
	index := -1

	if low < high {
		q = partition(arr, low, high)
		length := q - low + 1
		if length == n {
			index = q
		} else if length < n {
			index = FindNMax(arr, q+1, high, n-length)
		} else {
			index = FindNMax(arr, low, q-1, n)
		}
	}
	return index
}

func partition(arr []int, s, t int) (n int) {
	x := arr[s]
	i := s
	j := t
	for i < j {
		for arr[j] < x && i < j {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for arr[i] >= x && i < j {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = x
	n = i
	return
}
