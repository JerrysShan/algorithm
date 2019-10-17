package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	node2 := &ListNode{
		Val: 2,
	}

	node3 := &ListNode{
		Val: 3,
	}

	node4 := &ListNode{
		Val: 4,
	}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	head = swapPairs(head)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
