package GraphBFS

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/05-CC-BFS/Graph"
	"container/list"
)

type GraphBFS struct {
	g       *Graph.Graph
	visited []bool
	order  []int
}

func NewGraphBFS(g *Graph.Graph) (*GraphBFS) {
	gd := new(GraphBFS)
	gd.g = g
	gd.visited = make([]bool, g.V())
	for i := 0; i < g.V(); i++ {
		if !gd.visited[i] {
			gd.bfs(i)
		}
	}
	return gd
}

func (g *GraphBFS) bfs(v int) {
	queue := new(list.List)
	queue.PushBack(v)
	g.visited[v] = true
	for queue.Front() != nil {
		v := queue.Remove(queue.Front()).(int)
		g.order = append(g.order, v)
		for _, value := range g.g.Adj(v) {
			if !g.visited[value] {
				g.visited[value] = true
				queue.PushBack(value)
			}
		}
	}
}

func (g *GraphBFS) Order() []int {
	return g.order
}
