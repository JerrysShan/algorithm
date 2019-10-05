package leetcode

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	node2 := &ListNode{
		Val: 1,
	}

	node3 := &ListNode{
		Val: 3,
	}

	head.Next = node2
	node2.Next = node3
	deleteDuplicates(head)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
