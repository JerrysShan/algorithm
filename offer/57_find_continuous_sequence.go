package offer

func findContinuousSequence(target int) [][]int {
	if target < 2 {
		return nil
	}
	var ret [][]int
	for i := 1; i < target; i++ {
		temp := []int{i}
		num := i
		for j := i + 1; j < target; j++ {
			if num+j == target {
				temp = append(temp, j)
				ret = append(ret, temp)
				break
			}

			if num+j < target {
				num += j
				temp = append(temp, j)
			} else {
				break
			}

		}
	}
	return ret
}
