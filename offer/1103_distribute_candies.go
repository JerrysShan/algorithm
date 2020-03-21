package offer

func distributeCandies(candies int, num_people int) []int {
	arr := make([]int, num_people)
	i := 0
	for candies > 0 {
		index := i % num_people
		n := i / num_people
		if candies > n*num_people+index+1 {
			arr[index] = arr[index] + n*num_people + index + 1
			candies = candies - (n*num_people + index + 1)
		} else {
			arr[index] = arr[index] + candies
			candies = 0
		}
		i++
	}
	return arr
}
