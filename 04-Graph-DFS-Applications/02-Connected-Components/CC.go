package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/02-Connected-Components/Graph"
	"fmt"
)

type CC struct {
	g       *Graph.Graph
	visited []int
	cccount int
}

func NewCC(g *Graph.Graph) (*CC) {
	gd := new(CC)
	gd.g = g
	for i:=0;i< g.V();i++  {
		gd.visited = append(gd.visited, -1)
	}
	gd.Dfs()
	return gd
}

func (g *CC) Dfs() {
	for i := 0; i < g.g.V(); i++ {
		if g.visited[i] == -1 {
			g.dfs(i)
			g.cccount++
		}
	}
}

func (g *CC) dfs(v int) {
	g.visited[v] = g.cccount
	for _, value := range g.g.Adj(v) {
		if g.visited[value] == -1 {
			g.dfs(value)
		}
	}
}

func (g *CC) Count() int {
	for _,v :=range g.visited{
		fmt.Print(v," ")
	}
	fmt.Println()
	return g.cccount
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/02-Connected-Components/g.txt")
	cc := NewCC(g)
	fmt.Println(cc.Count())
}
