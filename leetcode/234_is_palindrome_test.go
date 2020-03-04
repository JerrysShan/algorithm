package leetcode

import "testing"

func TestIsPalindromeList(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	head.Next = node2

	node3 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 1,
	}
	node2.Next = node3
	node3.Next = node4
	t.Log(isPalindromeList(head))
}
