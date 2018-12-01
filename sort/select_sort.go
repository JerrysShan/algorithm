package sort

// func main() {
// 	arr := []int{1, 23, 4, 5, 2, 9, 10}
// 	selectSort(&arr)
// 	for _, item := range arr {
// 		fmt.Println(item)
// 	}
// }

func selectSort(arr *[]int) {
	length := len(*arr)
	var k int
	for i := 0; i < length-1; i++ {
		k = i
		for j := i + 1; j <= length-1; j++ {
			if (*arr)[j] < (*arr)[k] {
				k = j
			}
		}
		(*arr)[i], (*arr)[k] = (*arr)[k], (*arr)[i]
	}
}
