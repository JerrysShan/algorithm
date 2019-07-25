package btree

import (
	"algorithm/common"
	"algorithm/structure"
	"fmt"
)

// BinTreeMiddleOrderRecursion 二叉树的中序遍历递归算法
func BinTreeMiddleOrderRecursion(root *common.Node) {
	if root == nil {
		return
	}
	BinTreeMiddleOrderRecursion(root.Left)
	fmt.Print(root.Data, "->")
	BinTreeMiddleOrderRecursion(root.Right)
}

// BinTreeMiddleOrder 二叉树的中序遍历非递归算法
func BinTreeMiddleOrder(root *common.Node) {
	if root == nil {
		return
	}
	stack := &structure.Stack{}
	cur := root
	for cur != nil || stack.Size() > 0 {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		temp := stack.Pop().(*common.Node)
		fmt.Println(temp.Data, "->")
		cur = temp.Right
	}
}
