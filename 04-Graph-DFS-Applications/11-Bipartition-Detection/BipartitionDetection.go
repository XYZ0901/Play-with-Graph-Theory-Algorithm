package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/11-Bipartition-Detection/Graph"
	"fmt"
)

type BipartitionDetection struct {
	g           *Graph.Graph
	visited     []bool
	colors      []int
	isBipartite bool
}

func NewBipartitionDetection(g *Graph.Graph) (*BipartitionDetection) {
	gd := new(BipartitionDetection)
	gd.g = g
	gd.isBipartite = true
	gd.colors = make([]int, g.V())
	for i := 0; i < g.V(); i++ {
		gd.colors[i] = -1
	}
	gd.visited = make([]bool, g.V())

	for i := 0; i < g.V(); i++ {
		if !gd.visited[i] {
			if !gd.dfs(i, 0) {
				gd.isBipartite = false
				break
			}
		}
	}
	return gd
}

func (g *BipartitionDetection) dfs(v, color int) bool {
	g.visited[v] = true
	g.colors[v] = color
	for _, value := range g.g.Adj(v) {
		if !g.visited[value] {
			if !g.dfs(value, 1-color) {
				return false
			}
		} else if g.colors[value] == g.colors[v] {
			return false
		}
	}
	return true
}

func (g *BipartitionDetection) IsBipartite() bool {
	return g.isBipartite
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/11-Bipartition-Detection/g.txt")
	bipartitionDetection := NewBipartitionDetection(g)
	fmt.Println(bipartitionDetection.IsBipartite())

	g2 := Graph.NewGraph("./04-Graph-DFS-Applications/11-Bipartition-Detection/g2.txt")
	bipartitionDetection2 := NewBipartitionDetection(g2)
	fmt.Println(bipartitionDetection2.IsBipartite())
}
