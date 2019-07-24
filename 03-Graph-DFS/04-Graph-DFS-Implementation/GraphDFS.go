package main

import (
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/04-Graph-DFS-Implementation/Graph"
	"fmt"
)

type GraphDFS struct {
	g       *Graph.Graph
	visited []bool
	order   []int
}

func NewGraphDFS(g *Graph.Graph) (*GraphDFS) {
	gd := new(GraphDFS)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.Dfs()
	return gd
}

func (g *GraphDFS) Dfs() {
	g.dfs(0, g.visited, &g.order)
}

func (g *GraphDFS) dfs(v int, visited []bool, list *[]int) {
	visited[v] = true
	*list = append(*list, v)
	adj := g.g.Adj(v)
	a := adj.Traverse()
	for _, value := range a {
		if !visited[value] {
			g.dfs(value, visited, list)
		}
	}
}

func main() {
	g := Graph.NewGraph("./03-Graph-DFS/04-Graph-DFS-Implementation/g.txt")
	graphDFS := NewGraphDFS(g)
	fmt.Println(graphDFS.order)
}
