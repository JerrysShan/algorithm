package leetcode

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[i]标识凑零钱i的最小值
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func split(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	result := -1
	for i := 0; i < len(coins); i++ {
		min := split(coins, amount-coins[i])
		if min == -1 {
			continue
		}
		if result == -1 {
			result = min
		} else if result > min+1 {
			result = min + 1
		}
	}
	return result
}
