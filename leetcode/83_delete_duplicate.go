package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[int]bool)
	p := head
	n := p.Next
	m[p.Val] = true
	for n != nil {
		if _, ok := m[n.Val]; ok {
			p.Next = n.Next
		} else {
			m[n.Val] = true
			p = n
		}
		n = n.Next
	}
	return head
}
