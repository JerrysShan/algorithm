package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	ret := 0
	diameter(root, &ret)
	return ret - 1
}

func diameter(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}
	l := diameter(root.Left, ret)
	r := diameter(root.Right, ret)
	if l+r+1 > *ret {
		*ret = l + r + 1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
