package leetcode

func maxProfit2(prices []int) int {
	price := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			price += prices[i] - prices[i-1]
		}
	}
	return price
}
