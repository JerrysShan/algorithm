package leetcode

import "fmt"

func subsets(nums []int) [][]int {
	ret := [][]int{}
	tmp := make([]int, 0)
	helper(0, nums, &ret, &tmp)
	return ret
}

func helper(i int, nums []int, res *[][]int, tmp *[]int) {
	dst := make([]int, len(*tmp))
	copy(dst, *tmp)
	*res = append(*res, dst)
	for j := i; j < len(nums); j++ {
		*tmp = append(*tmp, nums[j])
		fmt.Println(*tmp, "before")
		helper(j+1, nums, res, tmp)
		fmt.Println(*tmp, "after")
		*tmp = (*tmp)[:len(*tmp)-1]
	}
}
