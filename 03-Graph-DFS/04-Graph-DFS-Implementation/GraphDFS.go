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
	g.dfs(0)
}

func (g *GraphDFS) dfs(v int) {
	g.visited[v] = true
	g.order = append(g.order, v)
	adj := g.g.Adj(v)

	for _, value := range adj {
		if !g.visited[value] {
			g.dfs(value)
		}
	}
}

func main() {
	g := Graph.NewGraph("./03-Graph-DFS/04-Graph-DFS-Implementation/g.txt")
	graphDFS := NewGraphDFS(g)
	fmt.Println(graphDFS.order)
}
