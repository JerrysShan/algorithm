package loadbalance

import (
	"strconv"
	"testing"
)

func TestHashCode(t *testing.T) {
	t.Log(hashCode("192.168.1.1"))
	t.Log(hashCode("192.168.1.2"))
	t.Log(hashCode("192.168.1.3"))
	t.Log(hashCode("192.168.1.1"))
}

func TestHashGet(t *testing.T) {
	hash := NewHash(20)
	hash.Add(arr...)
	for i := 0; i < 20; i++ {
		t.Log(hash.Get("192.168.1.19" + strconv.Itoa(i)))
		t.Log(hash.Get("192.168.1.19" + strconv.Itoa(i)))
	}
}
