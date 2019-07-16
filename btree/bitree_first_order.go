package btree

import (
	"algorithm/common"
	"algorithm/structure"
	"fmt"
)

// BinTreeFirstOrderRecursion 二叉树的先序遍历递归算法
func BinTreeFirstOrderRecursion(root *common.Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, "->")
	BinTreeFirstOrderRecursion(root.Left)
	BinTreeFirstOrderRecursion(root.Right)
}

// BinTreeFirstOrder 二叉树的先序遍历非递归算法
func BinTreeFirstOrder(root *common.Node) {
	if root == nil {
		return
	}
	stack := &structure.Stack{}
	cur := root
	for stack.Size() > 0 || cur != nil {
		if cur != nil {
			fmt.Print(cur.Data, "->")
			stack.Push(cur)
			cur = cur.Left
		} else {
			top := stack.Top().(common.Node)
			cur = top.Right
			stack.Pop()
		}
	}
}
