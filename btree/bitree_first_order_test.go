package btree

import (
	"algorithm/common"
	"testing"
)

var node1 = &common.Node{Data: 1}
var node2 = &common.Node{Data: 2}
var node3 = &common.Node{Data: 3}
var node4 = &common.Node{Data: 4}
var node5 = &common.Node{Data: 5}
var node6 = &common.Node{Data: 6}
var node7 = &common.Node{Data: 7}

func TestBinTreeFirstOrder(t *testing.T) {
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Left = node6
	node5.Right = node7
	BinTreeFirstOrder(node1)
}

func TestBinTreeFirstOrderRecursion(t *testing.T) {
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Left = node6
	node5.Right = node7
	BinTreeFirstOrderRecursion(node1)
}
