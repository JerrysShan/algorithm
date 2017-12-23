package graph

import (
	"container/list"
	"fmt"
	"strconv"
)

// func main() {
// 	graph := []string{"v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"}
// 	arr := [][]int{
// 		{0, 1, 0, 1, 0, 0, 0, 0},
// 		{1, 0, 1, 0, 1, 0, 0, 0},
// 		{0, 1, 0, 0, 0, 0, 0, 1},
// 		{1, 0, 0, 0, 0, 1, 1, 0},
// 		{0, 1, 0, 0, 0, 0, 0, 1},
// 		{0, 0, 0, 1, 0, 0, 1, 0},
// 		{0, 0, 0, 1, 0, 1, 0, 0},
// 		{0, 0, 1, 0, 1, 0, 0, 0},
// 	}

// 	BFS(arr, graph)

// 	// for _, val := range graph {
// 	// 	fmt.Println(val)
// 	// }

// }

//广度优先搜索
func BFS(arr [][]int, graph []string) {
	var visited [8]bool
	for i := 0; i < 8; i++ {
		visited[i] = false
	}
	queue := list.New()
	queue.PushBack(0)
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		str := fmt.Sprint(element.Value)
		index, _ := strconv.Atoi(str)
		if !visited[index] {
			visited[index] = true
			fmt.Print(graph[index] + "->")
			for i := 0; i < len(arr[index]); i++ {
				if arr[index][i] == 1 {
					queue.PushBack(i)
				}
			}
		}
	}
}
