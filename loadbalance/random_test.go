package loadbalance

import "testing"

var arr = []string{"http://192.168.192.1", "http://192.168.192.2", "http://192.168.192.3"}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		ret := Random(arr)
		t.Log(ret)
	}
}

func TestWeightRandom(t *testing.T) {
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
	for i := 0; i < 10; i++ {
		host := WeightRandom(hosts)
		t.Log(host)
	}
}
