package offer

import "testing"

func TestKnapsack(t *testing.T) {
	t.Log(knapsack(150, []int{35, 30, 60, 50, 40, 10, 25}, []int{10, 40, 30, 50, 35, 40, 30}))
	t.Log(knapsack(30, []int{28, 12, 12}, []int{30, 20, 20}))
	t.Log(knapsack2(150, []int{35, 30, 60, 50, 40, 10, 25}, []int{10, 40, 30, 50, 35, 40, 30}))
	t.Log(knapsack2(30, []int{28, 12, 12}, []int{30, 20, 20}))
}
