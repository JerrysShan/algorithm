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

func calWays2(n int) int {
	if n == 1 {
		return 1
	}
	a := 1
	b := 1
	temp := a + b
	for i := 2; i <= n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return temp
}
