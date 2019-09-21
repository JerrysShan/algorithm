package btree

import "algorithm/common"

func leftRotation(x *common.Node) {
	y := x.Right
	t2 := y.Left
	y.Left = x
	x.Right = t2
	// 更新高度
}

func rightRotation(x *common.Node) {
	y := x.Left
	t2 := y.Right
	y.Right = x
	x.Left = t2
	// 更新高度
}
