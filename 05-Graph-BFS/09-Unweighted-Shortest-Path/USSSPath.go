package main

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/09-Unweighted-Shortest-Path/Graph"
	"container/list"
	"fmt"
)

// Unweighted Single Source Shortest Path
type USSSPath struct {
	g *Graph.Graph
	s int

	visited []bool
	pre     []int
	dis     []int
}

func NewUSSSPath(g *Graph.Graph, s int) *USSSPath {

	this := new(USSSPath)
	this.g = g
	this.s = s

	this.visited = make([]bool, g.V())
	this.pre, this.dis = make([]int, g.V()), make([]int, g.V())
	for k, _ := range this.pre {
		this.pre[k] = -1
		this.dis[k] = -1
	}

	this.bfs()

	for _, v := range this.dis {
		fmt.Print(v, " ")
	}
	fmt.Println()

	return this
}

func (g *USSSPath) bfs() {
	s := g.s
	queue := new(list.List)
	queue.PushBack(s)
	g.visited[s] = true
	g.pre[s] = s
	g.dis[s] = 0

	for queue.Front() != nil {
		v := queue.Remove(queue.Front()).(int)
		for _, value := range g.g.Adj(v) {
			if !g.visited[value] {
				g.visited[value] = true
				queue.PushBack(value)
				g.pre[value] = v
				g.dis[value] = g.dis[v] + 1
			}
		}
	}
}

func (g *USSSPath) IsConnectedTo(t int) bool {
	g.g.ValidateVertex(t)
	return g.visited[t]
}

func (g *USSSPath) Path(t int) []int {
	res := []int{}
	if !g.IsConnectedTo(t) {
		return res
	}

	cur := t
	for cur != g.s {
		res = append(res, cur)
		cur = g.pre[cur]
	}
	res = append(res, g.s)

	reverse(&res)

	return res
}

func (g *USSSPath) Dis(t int) int {
	g.g.ValidateVertex(t)
	return g.dis[t]
}

func main() {
	g := Graph.NewGraph("./05-Graph-BFS/09-Unweighted-Shortest-Path/g.txt")
	ussspath := NewUSSSPath(g, 0)
	fmt.Println("0 -> 6 : ", ussspath.Path(6))
	fmt.Println("0 -> 6 : ", ussspath.Dis(6))
}

func reverse(res *[]int) {
	for i := 0; i < len(*res)/2; i++ {
		temp := (*res)[i]
		(*res)[i] = (*res)[len(*res)-i-1]
		(*res)[len(*res)-i-1] = temp
	}
}
