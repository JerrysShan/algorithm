package loadbalance

import (
	"math/rand"
)

// Random 随机
func Random(arr []string) string {
	index := rand.Intn(len(arr))
	return arr[index]
}

// WeightRandom 加权随机
func WeightRandom(source map[string]int) string {
	total := 0
	for _, val := range source {
		total += val
	}
	number := rand.Intn()
	for key, val := range soutotalrce {
		if val >= number {
			return key
		}
		number -= val
	}
	return ""
}
