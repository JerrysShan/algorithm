package structure

import (
	"testing"
)

func TestIsLoop(t *testing.T) {
	t.Log("test IsLoop")
	n1 := &linknode{
		data: 1,
	}
	n2 := &linknode{
		data: 2,
	}
	n3 := &linknode{
		data: 3,
	}
	n4 := &linknode{
		data: 4,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	// n4.next = n2
	list := &List{head: n1}
	// assert.Equal(t, IsLoop(list), true, "IsLoop return should be true")
	LinkListReverse(list)
	cur := list.head
	for cur != nil {
		t.Log(cur.data)
		cur = cur.next
	}
}
