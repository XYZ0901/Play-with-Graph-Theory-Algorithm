package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/11-Bipartition-Detection/Graph"
	"fmt"
)

type GraphDFS struct {
	g       *Graph.Graph
	visited []bool
	pre     []int
	post    []int
}

func NewGraphDFS(g *Graph.Graph) (*GraphDFS) {
	gd := new(GraphDFS)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.Dfs()
	return gd
}

func (g *GraphDFS) Dfs() {
	for i := 0; i < g.g.V(); i++ {
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

func (g *GraphDFS) dfs(v int) {
	g.visited[v] = true
	adj := g.g.Adj(v)
	g.pre = append(g.pre, v) // 先序
	for _, value := range adj {
		if !g.visited[value] {
			g.dfs(value)
		}
	}
	g.post = append(g.post, v) // 后序
}

func (g *GraphDFS) Pre() []int {
	return g.pre
}

func (g *GraphDFS) Post() []int {
	return g.post
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/09-Cycle-Detection/g.txt")
	graphDFS := NewGraphDFS(g)
	fmt.Println(graphDFS.Pre())
	fmt.Println(graphDFS.Post())
}
