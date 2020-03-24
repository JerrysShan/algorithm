package offer

func kthLargest(root *TreeNode, k int) int {
	arr := []int{}
	mid(root, &arr)
	return arr[len(arr)-k]
}

func mid(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	mid(root.Left, arr)
	*arr = append(*arr, root.Val)
	mid(root.Right, arr)
}
