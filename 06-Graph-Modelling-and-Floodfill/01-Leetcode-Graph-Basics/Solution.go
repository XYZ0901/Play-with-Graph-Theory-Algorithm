package main

import "fmt"

var (
	visited []bool
	colors  []int
	Graph   [][]int
)

func isBipartite(graph [][]int) bool {
	var V = len(graph)
	Graph = graph
	visited = make([]bool, V)
	colors = make([]int, V)
	for v := 0; v < V; v++ {
		if !visited[v] {
			if !dfs(v, 0) {
				return false
			}
		}
	}
	return true
}

func dfs(v, color int) bool {
	visited[v] = true
	colors[v] = color
	for _, w := range Graph[v] {
		if !visited[w] {
			if !dfs(w, 1-color) {
				return false
			}
		} else if colors[v] == colors[w] {
			return false
		}
	}
	return true
}

func main() {
	graph1 := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	graph2 := [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}
	fmt.Println(isBipartite(graph1))
	fmt.Println(isBipartite(graph2))
}
