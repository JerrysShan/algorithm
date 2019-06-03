package loadbalance

import (
	"testing"
)

func TestRoundRobin(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret := RoundRobin(arr)
		t.Log(ret)
	}
}

func TestWeightRoundRobin(t *testing.T) {
	hosts := make(map[string]int)
	w1 := 1
	w2 := 3
	w3 := 6
	for index, v := range arr {
		switch index {
		case 0:
			hosts[v] = w1
		case 1:
			hosts[v] = w2
		case 2:
			hosts[v] = w3
		}
	}
	w := NewWeightRoundRobin(hosts)
	for i := 0; i < 10; i++ {
		t.Log(w.Next())
	}
}
