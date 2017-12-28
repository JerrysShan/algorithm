package structure

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestIsLoop(t *testing.T) {
	t.Log("test IsLoop")
	n1 := &Node{
		data: 1,
	}
	n2 := &Node{
		data: 2,
	}
	n3 := &Node{
		data: 3,
	}
	n4 := &Node{
		data: 4,
	}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n2
	assert.Equal(t, IsLoop(n1), true, "IsLoop return should be true")
}
