package leetcode

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	for i := m; i < m+n+1; i++ {
		for j := i - 1; j > 0; j-- {
			if nums1[j] < nums1[j-1] {
				nums1[j], nums1[j-1] = nums1[j-1], nums1[j]
			}
		}
	}
	fmt.Println(nums1)
}
