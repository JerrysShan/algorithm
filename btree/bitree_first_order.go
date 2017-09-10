package btree

import (
	"algorithm/common"
	"fmt"
)

//二叉树的先序遍历递归算法
func BinTreeFirstOrderRecursion(root *common.Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Data, "->")
	BinTreeFirstOrderRecursion(root.Left)
	BinTreeFirstOrderRecursion(root.Right)
}

//二叉树的先序遍历非递归算法

func BinTreeFirstOrder() {

}
