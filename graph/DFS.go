package main

import "fmt"

var visited [8]bool

func main() {
	graph := []string{"v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"}
	arr := [][]int{
		{0, 1, 0, 1, 0, 0, 0, 0},
		{1, 0, 1, 0, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0},
	}
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			DFS(arr, graph, i)
		}
	}
}

//深度优先遍历
func DFS(arr [][]int, graph []string, index int) {
	fmt.Print(graph[index] + "->")
	visited[index] = true
	for i := 0; i < len(arr[index]); i++ {
		if arr[index][i] == 1 {
			if !visited[i] {
				DFS(arr, graph, i)
			}
		}
	}
}
