package sort

func SimpleCountSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	min := arr[0]
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	countArr := make([]int, max-min+1)
	for _, ele := range arr {
		countArr[ele-min]++
	}
	index := 0
	for k, ele := range countArr {
		for i := 0; i < ele; i++ {
			arr[index] = k + min
			index++
		}
	}
}

func SuperCountSort() {

}
