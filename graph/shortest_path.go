package graph

import (
	"fmt"
	"math"
)

var MAXNUM = math.MaxInt32

func ShortestPathDIJ(g [][]int) {
	p := make([][]bool, len(g)) //p[v][w] 为true,则w是从v0到v当前求得最短路径上的顶点
	for k := 0; k < len(g); k++ {
		p[k] = make([]bool, len(g))
	}
	final := make([]bool, len(g))
	d := make([]int, len(g))
	path := make([]int, len(g))
	v0 := 0
	path[v0] = 0
	for v := 0; v < len(g); v++ {
		final[v] = false
		d[v] = g[v0][v]
		for w := 0; w < len(g); w++ {
			p[v][w] = false
		}
		if d[v] < MAXNUM {
			p[v][v0] = true
			p[v][v] = true
			path[v] = v0
		} else {
			path[v] = -1
		}
	}
	d[v0] = 0
	final[v0] = true
	for i := 1; i < len(g); i++ {
		min := MAXNUM
		var v int
		for w := 0; w < len(g); w++ {
			if !final[w] {
				if d[w] < min {
					v = w
					min = d[w]
				}
			}
		}
		final[v] = true
		// 跟新当前最短路径及其距离
		for w := 0; w < len(g); w++ {
			if (!final[w]) && (min+g[v][w] < d[w]) {
				d[w] = min + g[v][w]
				fmt.Println("d[w]", d[w])
				p[w] = p[v]
				p[w][w] = true
				path[w] = v
			}
		}
	}
	for k, val := range d {
		fmt.Println("short path:", k, val)
	}
	for index, value := range p {
		fmt.Println(index, value)
	}
	for k1, v1 := range path {
		fmt.Println(k1, v1)
	}
}
