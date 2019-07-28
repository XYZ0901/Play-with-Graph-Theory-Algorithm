package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/05-Single-Source-Path/Graph"
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
	for i := 0; i < g.V(); i++ {
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
	return g.cccount
}

func (g *CC) IsConnected(v, w int) bool {
	g.g.ValidateVertex(v)
	g.g.ValidateVertex(w)

	return g.visited[v] == g.visited[w]
}

func (g *CC) Components() [][]int {
	res := [][]int{}

	for i := 0; i < g.cccount; i++ {
		tmp := make([]int, 0)
		res = append(res, tmp)
	}

	for v := 0; v < g.g.V(); v++ {
		res[g.visited[v]] = append(res[g.visited[v]], v)
	}

	return res
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/05-Single-Source-Path/g.txt")
	cc := NewCC(g)
	fmt.Println(cc.Count())
	fmt.Println(cc.IsConnected(0, 6))
	fmt.Println(cc.IsConnected(0, 5))

	comp := cc.Components()
	for ccid:=0;ccid< len(comp);ccid++  {
		fmt.Print(ccid," : ")
		for _,w := range comp[ccid]{
			fmt.Print(w," ")
		}
		fmt.Println()
	}
}
