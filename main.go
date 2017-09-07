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
	bs.Put(9)
	for _, val := range bs.Keys {
		fmt.Print(val, " ")
	}
	fmt.Println("\n")
	fmt.Println(bs.Rank(6))
	fmt.Println(bs.Ceiling(6))
	fmt.Println(bs.Floor(6))
	fmt.Println(bs.Floor(11))

	fmt.Println(bs.Floor(1))
	fmt.Println(bs.Get(1))
}
