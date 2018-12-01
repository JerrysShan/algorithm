package dynamicprogram

func ClimbStair(n int) int {
	return calcWays(n)
}

func calcWays(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return calcWays(n-1) + calcWays(n-2)
}
