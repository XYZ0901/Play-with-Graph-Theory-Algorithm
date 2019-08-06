package main

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/07-Bipartition-Detection-BFS/Graph"
	"container/list"
	"fmt"
)

type BipartitionDetection struct {
	g       *Graph.Graph
	visited []bool
	color  []int
	isBipartite bool
}

func NewBipartitionDetection(g *Graph.Graph) (*BipartitionDetection) {
	gd := new(BipartitionDetection)
	gd.g = g
	gd.isBipartite=true
	gd.visited = make([]bool, g.V())
	gd.color = make([]int,g.V())
	for k,_ := range gd.color{
		gd.color[k] = -1
	}

	for i := 0; i < g.V(); i++ {
		if !gd.visited[i] {
			if !gd.bfs(i){
				gd.isBipartite = false
				break
			}
		}
	}
	return gd
}

func (g *BipartitionDetection) bfs(v int) bool {
	queue := new(list.List)
	queue.PushBack(v)
	g.visited[v] = true
	g.color[v] = 0
	for queue.Front() != nil {
		v := queue.Remove(queue.Front()).(int)
		for _, value := range g.g.Adj(v) {
			if !g.visited[value]{
				g.visited[value] = true
				g.color[value] = 1-g.color[v]
				queue.PushBack(value)
			}else if g.color[v] == g.color[value] {
				return false
			}
		}
	}
	return true
}

func (g *BipartitionDetection)IsBipartite() bool {
	return g.isBipartite
}

func main() {
	g := Graph.NewGraph("./05-Graph-BFS/07-Bipartition-Detection-BFS/g.txt")
	bipartitionDetection := NewBipartitionDetection(g)
	fmt.Println(bipartitionDetection.IsBipartite())

	g2 := Graph.NewGraph("./05-Graph-BFS/07-Bipartition-Detection-BFS/g2.txt")
	bipartitionDetection2 := NewBipartitionDetection(g2)
	fmt.Println(bipartitionDetection2.IsBipartite())

	g3 := Graph.NewGraph("./05-Graph-BFS/07-Bipartition-Detection-BFS/g3.txt")
	bipartitionDetection3 := NewBipartitionDetection(g3)
	fmt.Println(bipartitionDetection3.IsBipartite())
}
