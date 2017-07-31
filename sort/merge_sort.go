package main

func main() {
	arr := []int{1, 23, 4, 5, 2, 9, 10}
	result := mergeSort(arr)
	for _, item := range result {
		fmt.Println(item)
	}

}

//二路归并排序
func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
