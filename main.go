package main

import (
	"algorithm/search"
	"fmt"
)

func main() {
	bs := &search.BinarySearch{
		Keys: []int{2, 3, 4, 5, 8, 10, 0},
		N:    6,
	}
	bs.Put(1)
	for _, val := range bs.Keys {
		fmt.Println(val)
	}
	fmt.Println(search.Total)
}
