package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	head = head.Next
	var prev *ListNode
	for p != nil && p.Next != nil {
		next := p.Next
		nextnext := next.Next
		next.Next = p
		p.Next = nextnext
		if prev != nil {
			prev.Next = next
		}
		prev = p
		p = nextnext
	}

	return head
}
