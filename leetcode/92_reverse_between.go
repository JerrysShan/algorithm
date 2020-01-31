package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	index := 1
	temp := head

	var start, prev *ListNode
	if index == m {
		prev = nil
		start = head
	}

	for temp != nil {
		if index+1 == m {
			prev = temp
			start = temp.Next
		}
		index++
		temp = temp.Next
	}
	var nextnext *ListNode
	cur := start
	next := start.Next
	for n-m != 0 {
		nextnext = next.Next
		next.Next = cur
		cur = next
		next = nextnext
		n--
	}
	start.Next = next
	if prev == nil {
		head = cur
	} else {
		prev.Next = cur
	}

	return head
}
