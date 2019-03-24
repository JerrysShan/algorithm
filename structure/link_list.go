package structure

type linknode struct {
	data interface{}
	next *linknode
}

type List struct {
	head *linknode
}

func newList() *List {
	return &List{head: &linknode{next: nil}}
}

// IsLoop 判断一个链表是否有环
func IsLoop(list *List) bool {
	head := list.head
	if head == nil {
		return false
	}
	m := head
	n := head.next.next
	for {

		if m == nil || n == nil {
			break
		}
		if m == n {
			return true
		}
		m = m.next
		n = n.next.next
	}
	return false
}

// IsCross 判断两个链表是否有交叉
func IsCross(listA, listB *List) bool {
	if listA.head == nil || listB.head == nil {
		return false
	}
	curA := listA.head
	curB := listB.head
	for curA.next != nil {
		curA = curA.next
	}
	for curB.next != nil {
		curB = curB.next
	}
	if curA == curB {
		return true
	}
	return false
}

// FindCross 找到两个链表的交叉处
func FindCross(listA, listB *List) {
	if !IsCross(listA, listB) {

	}

}

// LinkListReverse 单向链表的翻转
func LinkListReverse(list *List) *List {
	if list == nil {
		return list
	}
	cur := list.head
	next := cur.next
	for next != nil {
		nextnext := next.next
		next.next = cur
		cur = next
		next = nextnext
	}
	return list
}
