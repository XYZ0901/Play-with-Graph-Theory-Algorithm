package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/05-Single-Source-Path/Graph"
	"fmt"
)

type SingleSourcePath struct {
	g *Graph.Graph
	s int

	visited []bool
	pre     []int
}

func NewGraphDFS(g *Graph.Graph, s int) (*SingleSourcePath) {
	gd := new(SingleSourcePath)
	g.ValidateVertex(s)
	gd.g = g
	gd.s = s
	gd.visited = make([]bool, g.V())
	gd.pre = make([]int, g.V())
	for i := 0; i < g.V(); i++ {
		gd.pre[i] = -1
	}

	gd.Dfs()

	return gd
}

func (g *SingleSourcePath) Dfs() {
	g.dfs(g.s)
}

func (g *SingleSourcePath) dfs(v int) {
	g.visited[v] = true
	for _, value := range g.g.Adj(v) {
		if !g.visited[value] {
			g.pre[value] = v
			g.dfs(value)
		}
	}
}

func (singleSourcePath *SingleSourcePath) IsConnectedTo(t int) bool {
	singleSourcePath.g.ValidateVertex(t);
	return singleSourcePath.visited[t]
}

func (singleSourcePath *SingleSourcePath) Path(t int) []int {
	res := []int{}
	if !singleSourcePath.IsConnectedTo(t) {
		return res
	}
	cur := t
	for cur != -1{
		res = append(res,cur)
		cur = singleSourcePath.pre[cur]
	}
	reverse(&res)
	return res
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/05-Single-Source-Path/g.txt")
	sspath := NewGraphDFS(g,0)
	fmt.Println("0 -> 6 : ",sspath.Path(6))
	fmt.Println("0 -> 5 : ",sspath.Path(5))
}

func reverse(res *[]int) {
	for i:=0;i<len(*res)/2;i++ {
		temp := (*res)[i]
		(*res)[i] = (*res)[len(*res)-i-1]
		(*res)[len(*res)-i-1] = temp
	}
}