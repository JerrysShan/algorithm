package loadbalance

import (
	"fmt"
	"sync"
)

var index = 0

// RoundRobin 轮询
func RoundRobin(arr []string) string {
	if index == len(arr) {
		index = 0
	}
	result := arr[index]
	index++
	return result
}

type host struct {
	Name          string
	CurrentWeight int
	Weight        int
}

// WeightRoundRobin 加权轮询
type WeightRoundRobin struct {
	Hosts       []*host
	TotalWeight int
	sync.Mutex
}

// NewWeightRoundRobin init WeightRoundRobin
func NewWeightRoundRobin(hosts map[string]int) *WeightRoundRobin {
	w := new(WeightRoundRobin)
	for k, v := range hosts {
		w.TotalWeight += v
		w.Hosts = append(w.Hosts, &host{Name: k, CurrentWeight: 0, Weight: v})
	}
	return w
}

// Next 返回一个可用实例
func (w *WeightRoundRobin) Next() string {
	allWeight := 0
	index := -1
	size := len(w.Hosts)
	w.Lock()
	for i := 0; i < size; i++ {
		w.Hosts[i].CurrentWeight += w.Hosts[i].Weight
		allWeight += w.Hosts[i].Weight
		if index == -1 || w.Hosts[index].CurrentWeight < w.Hosts[i].CurrentWeight {
			index = i
		}
	}
	w.Hosts[index].CurrentWeight -= allWeight
	w.Unlock()
	fmt.Println(index)
	for _, v := range w.Hosts {
		fmt.Println(v.Name, v.CurrentWeight, v.Weight)
	}
	return w.Hosts[index].Name
}

// 2
// http://192.168.192.1 1 1
// http://192.168.192.2 3 3
// http://192.168.192.3 -4 6
// 1
// http://192.168.192.1 2 1
// http://192.168.192.2 -4 3
// http://192.168.192.3 2 6
// 2
// http://192.168.192.1 3 1
// http://192.168.192.2 -1 3
// http://192.168.192.3 -2 6
// 0
// http://192.168.192.1 -6 1
// http://192.168.192.2 2 3
// http://192.168.192.3 4 6
// 2
// http://192.168.192.1 -5 1
// http://192.168.192.2 5 3
// http://192.168.192.3 0 6
// 1
// http://192.168.192.1 -4 1
// http://192.168.192.2 -2 3
// http://192.168.192.3 6 6
// 2
// http://192.168.192.1 -3 1
// http://192.168.192.2 1 3
// http://192.168.192.3 2 6
// 2
// http://192.168.192.1 -2 1
// http://192.168.192.2 4 3
// http://192.168.192.3 -2 6
// 1
// http://192.168.192.1 -1 1
// http://192.168.192.2 -3 3
// http://192.168.192.3 4 6
// 2
// http://192.168.192.1 0 1
// http://192.168.192.2 0 3
// http://192.168.192.3 0 6
