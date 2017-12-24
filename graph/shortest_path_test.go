package graph

import "testing"

func TestShortestPathDIJ(t *testing.T) {
	var graph = [][]int{
		{MAXNUM, MAXNUM, 10, MAXNUM, 30, 100},
		{MAXNUM, MAXNUM, 5, MAXNUM, MAXNUM, MAXNUM},
		{MAXNUM, MAXNUM, MAXNUM, 50, MAXNUM, MAXNUM},
		{MAXNUM, MAXNUM, MAXNUM, MAXNUM, MAXNUM, 10},
		{MAXNUM, MAXNUM, MAXNUM, 20, MAXNUM, 60},
		{MAXNUM, MAXNUM, MAXNUM, MAXNUM, MAXNUM, MAXNUM},
	}
	ShortestPathDIJ(graph)
}

func TestShortestPahtFLOYD(t *testing.T) {
	var graph = [][]int{
		{0, 4, 11},
		{6, 0, 2},
		{3, MAXNUM, 0},
	}
	ShortestPathFLOYD(graph)
}
