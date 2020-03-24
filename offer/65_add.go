package offer

func add(a, b int) int {
	for b != 0 {
		plus := a ^ b
		b = (a & b) << 1
		a = plus
	}
	return a
}
