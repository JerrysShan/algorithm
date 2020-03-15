package offer

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	ret := [][]int{}
	arr := []int{}
	dfs2(root, sum, arr, &ret)
	return ret
}

func dfs2(root *TreeNode, sum int, arr []int, ret *([][]int)) {
	if root != nil {
		arr = append(arr, root.Val)
	}
	if root.Left == nil && root.Right == nil {
		s := 0
		for _, v := range arr {
			s += v
		}
		if s == sum {
			temp := make([]int, len(arr))
			copy(temp, arr)
			*ret = append(*ret, temp)
		}
		return
	}
	if root.Left != nil {
		dfs2(root.Left, sum, arr, ret)
	}
	if root.Right != nil {
		dfs2(root.Right, sum, arr, ret)
	}
}
