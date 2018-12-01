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
