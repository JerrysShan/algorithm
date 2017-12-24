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
		// 更新当前最短路径及其距离
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

func ShortestPathFLOYD(g [][]int) {
	fmt.Println("----------------FLOYD-----------------")
	p := make([][][]bool, len(g)) // p[v][w][u]
	d := make([][]int, len(g))
	for i := 0; i < len(g); i++ {
		d[i] = make([]int, len(g))
		p[i] = make([][]bool, len(g))
		for j := 0; j < len(g); j++ {
			p[i][j] = make([]bool, len(g))
		}
	}
	for v := 0; v < len(g); v++ {
		for w := 0; w < len(g); w++ {
			d[v][w] = g[v][w]
			if d[v][w] < MAXNUM {
				p[v][w][v] = true
				p[v][w][w] = true
			}
		}
	}
	for u := 0; u < len(g); u++ {
		for v := 0; v < len(g); v++ {
			for w := 0; w < len(g); w++ {
				if d[v][u]+d[u][w] < d[v][w] {
					d[v][w] = d[v][u] + d[u][w]
					for i := 0; i < len(g); i++ {
						p[v][w][i] = p[v][u][i] || p[u][w][i]
					}
				}
			}
		}
	}
	fmt.Println("-------------D---------------")
	for index1, li := range d {
		for index, val := range li {
			fmt.Println(index1, index, val)
		}
	}
	for index1, l1 := range p {
		for index2, l2 := range l1 {
			for index3, val := range l2 {
				fmt.Println(index1, index2, index3, val)
			}
		}
	}
}
