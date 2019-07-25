package main

import (
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/08-Graph-Interface/AdjList"
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/08-Graph-Interface/AdjMatrix"
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/08-Graph-Interface/AdjSet"
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/08-Graph-Interface/Graph"
	"fmt"
)

type GraphDFS struct {
	g       Graph.Graph
	visited []bool
	pre     []int
	post    []int
}

func NewGraphDFS(g Graph.Graph) (*GraphDFS) {
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

	g.pre = append(g.pre, v) // 先序
	adj := g.g.Adj(v)
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

	g1 := AdjSet.NewAdjSet("./03-Graph-DFS/08-Graph-Interface/g.txt")
	graphDFS1 := NewGraphDFS(g1)
	fmt.Println("DFS preOrder : ", graphDFS1.Pre())
	fmt.Println("DFS postOrder : ", graphDFS1.Post())

	g2 := AdjList.NewAdjList("./03-Graph-DFS/08-Graph-Interface/g.txt")
	graphDFS2 := NewGraphDFS(g2)
	fmt.Println("DFS preOrder : ", graphDFS2.Pre())
	fmt.Println("DFS postOrder : ", graphDFS2.Post())

	g3 := AdjMatrix.NewAdjMatrix("./03-Graph-DFS/08-Graph-Interface/g.txt")
	graphDFS3 := NewGraphDFS(g3)
	fmt.Println("DFS preOrder : ", graphDFS3.Pre())
	fmt.Println("DFS postOrder : ", graphDFS3.Post())
}
