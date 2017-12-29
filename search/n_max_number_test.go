package search

import (
	"fmt"
	"testing"
)

func TestFindNMax(t *testing.T) {
	arr := []int{20, 100, 4, 2, 87, 9, 8, 5, 46, 26}
	n := 4
	t.Log("test find Nth max number")
	{
		FindNMax(arr, 0, len(arr)-1, n)
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	}
}
