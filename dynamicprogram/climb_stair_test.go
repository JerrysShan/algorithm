package dynamicprogram

import "testing"

func TestClimbStair(t *testing.T) {
	n1 := ClimbStair(1)
	n2 := ClimbStair(2)
	n3 := ClimbStair(3)
	n4 := ClimbStair(4)
	n5 := ClimbStair(5)
	t.Log(n1, n2, n3, n4, n5)
}

func TestCalways2(t *testing.T) {
	t.Log(calWays2(1))
	t.Log(calWays2(2))
	t.Log(calWays2(3))
	t.Log(calcWays(4))
	t.Log(calcWays(5))
}
