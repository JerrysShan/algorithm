/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, l, h int) *TreeNode {
	if l > h {
		return nil
	}
	m := l + (h-l)/2
	root := &TreeNode{
		Val: nums[m],
	}
	root.Left = bst(nums, l, m-1)
	root.Right = bst(nums, m+1, h)
	return root
}
