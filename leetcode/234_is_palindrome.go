package leetcode

func isPalindromeList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow := head
	fast := head
	cur := head
	var pre *ListNode
	for fast != nil && fast.Next != nil {
		cur = slow
		slow = slow.Next
		fast = fast.Next.Next
		cur.Next = pre
		pre = cur
	}
	if fast != nil {
		slow = slow.Next
	}
	for pre != nil && slow != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
