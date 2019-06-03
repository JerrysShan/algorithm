package loadbalance

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

// Host is a host
type Host struct {
	Name          string
	CurrentWeight int
	Weight        int
}

// WeightRoundRobin 加权轮询
type WeightRoundRobin struct {
	Hosts       []*Host
	TotalWeight int
}

// NewWeightRoundRobin init WeightRoundRobin
func NewWeightRoundRobin(hosts map[string]int) *WeightRoundRobin {
	w := new(WeightRoundRobin)
	for k, v := range hosts {
		w.TotalWeight += v
		w.Hosts = append(w.Hosts, &Host{Name: k, CurrentWeight: 0, Weight: v})
	}
	return w
}

// Next 返回一个host实例
func (w *WeightRoundRobin) Next() string {
	allWeight := 0
	index := -1
	size := len(w.Hosts)
	for i := 0; i < size; i++ {
		w.Hosts[i].CurrentWeight += w.Hosts[i].Weight
		allWeight += w.Hosts[i].Weight
		if index == -1 || w.Hosts[index].CurrentWeight < w.Hosts[i].CurrentWeight {
			index = i
		}
	}
	w.Hosts[index].CurrentWeight -= allWeight
	return w.Hosts[index].Name
}
