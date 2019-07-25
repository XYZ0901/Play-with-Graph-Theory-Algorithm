package main

import (
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/07-Adjacency-Matrix-DFS/AdjMatrix"
	"fmt"
)

type AdjMatrixDFS struct {
	g       *AdjMatrix.AdjMatrix
	visited []bool
	pre     []int
	post    []int
}

func NewAdjMatrixDFS(g *AdjMatrix.AdjMatrix) (*AdjMatrixDFS) {
	gd := new(AdjMatrixDFS)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.Dfs()
	return gd
}

func (g *AdjMatrixDFS) Dfs() {
	for i:=0;i<g.g.V();i++ {
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

func (g *AdjMatrixDFS) dfs(v int) {
	g.visited[v]=true

	g.pre = append(g.pre, v) // 先序
	for _, value := range g.g.Adj(v) {
		if !g.visited[value] {
			g.dfs(value)
		}
	}
	g.post = append(g.post, v) // 后序
}

func (g *AdjMatrixDFS) Pre() []int {
	return g.pre
}

func (g *AdjMatrixDFS) Post() []int {
	return g.post
}

func main() {
	g := AdjMatrix.NewAdjMatrix("./03-Graph-DFS/07-Adjacency-Matrix-DFS/g.txt")
	graphDFS := NewAdjMatrixDFS(g)
	fmt.Println(graphDFS.Pre())
	fmt.Println(graphDFS.Post())
}
