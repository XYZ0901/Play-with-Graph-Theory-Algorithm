package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/01-Connected-Components-Count/Graph"
	"fmt"
)

type CC struct {
	g       *Graph.Graph
	visited []bool
	cccount int
}

func NewCC(g *Graph.Graph) (*CC) {
	gd := new(CC)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.Dfs()
	return gd
}

func (g *CC) Dfs() {
	for i := 0; i < g.g.V(); i++ {
		if !g.visited[i] {
			g.dfs(i)
			g.cccount++
		}
	}
}

func (g *CC) dfs(v int) {
	g.visited[v] = true
	adj := g.g.Adj(v)
	for _, value := range adj {
		if !g.visited[value] {
			g.dfs(value)
		}
	}
}

func (g *CC) Count() int {
	return g.cccount
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/01-Connected-Components-Count/g.txt")
	cc := NewCC(g)
	fmt.Println(cc.Count())
}
