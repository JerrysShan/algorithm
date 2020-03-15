package offer

func knapsack(w int, weights, vals []int) int {
	n := len(weights)
	ret := make(map[int]int, w+1) // 重量不超过k的最大价值v
	for i := 0; i < n; i++ {
		for j := w; j > weights[i]; j-- {
			if ret[j] < ret[j-weights[i]]+vals[i] {
				ret[j] = ret[j-weights[i]] + vals[i]
			}
		}
	}

	return ret[w]
}

func knapsack2(w int, weights, vals []int) int {
	n := len(weights)
	ret := make([][]int, n) // i个物品的重量不超过j的最大价值
	for i := 0; i < n; i++ {
		ret[i] = make([]int, w+1)
		ret[i][0] = 0
	}
	for j := 1; j <= w; j++ {
		if j >= weights[0] {
			ret[0][j] = vals[0]
		} else {
			ret[0][j] = 0
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= w; j++ {
			if j < weights[i] {
				ret[i][j] = ret[i-1][j]
			} else {
				if ret[i-1][j] > ret[i-1][j-weights[i]]+vals[i] {
					ret[i][j] = ret[i-1][j]
				} else {
					ret[i][j] = ret[i-1][j-weights[i]] + vals[i]
				}
			}
		}
	}
	return ret[n-1][w]
}
