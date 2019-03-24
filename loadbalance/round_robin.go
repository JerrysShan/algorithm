package loadbalance

var index = 0

// RoundRobin 轮询
func RoundRobin(arr []interface{}) interface{} {
	if index == len(arr) {
		index = 0
	}
	result := arr[index]
	index++
	return result
}

// WeightRoundRobin 加权轮询
func WeightRoundRobin(map[string]int) string {

	return ""
}
