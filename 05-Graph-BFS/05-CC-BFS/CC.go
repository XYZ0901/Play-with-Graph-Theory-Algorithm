package main

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/05-CC-BFS/Graph"
	"container/list"
	"fmt"
)

type CC struct {
	g       *Graph.Graph
	visited []int
	cccount int
}

func NewCC(graph *Graph.Graph) *CC {
	cc := new(CC)
	cc.g = graph
	cc.cccount = 0

	cc.visited = make([]int, graph.V())
	for i := 0; i < graph.V(); i++ {
		cc.visited[i] = -1
	}

	for i:=0;i<graph.V() ; i++ {
		if cc.visited[i] == -1 {
			cc.bfs(i)
		}
	}

	return cc
}

func (g *CC) bfs(v int) {
	quene := new(list.List)
	quene.PushBack(v)
	g.visited[v] = g.cccount

	for quene.Front()!= nil {
		v:=quene.Remove(quene.Front()).(int)
		for _,value := range g.g.Adj(v){
			if g.visited[value] == -1 {
				g.visited[value] = g.cccount
				quene.PushBack(value)
			}
		}
	}

	g.cccount++
}

func (g *CC)Count() int {
	return g.cccount
}

func (g *CC)IsCounted(v,w int) bool {
	g.g.ValidateVertex(w)
	g.g.ValidateVertex(v)
	return g.visited[v] == g.visited[w]
}

func (g *CC)Components() [][]int {
	res := [][]int{}

	for i:=0;i<g.cccount ; i++ {
		res = append(res,make([]int,0))
	}

	for k,v := range g.visited  {
		res[v] = append(res[v], k)
	}

	return res
}

func main() {
	g := Graph.NewGraph("./05-Graph-BFS/05-CC-BFS/g.txt")
	cc := NewCC(g)
	fmt.Println(cc.Count())
	fmt.Println(cc.Components())
}
