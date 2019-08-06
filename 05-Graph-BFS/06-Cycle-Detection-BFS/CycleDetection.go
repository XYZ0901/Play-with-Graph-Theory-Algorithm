package main

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/06-Cycle-Detection-BFS/Graph"
	"container/list"
	"fmt"
)

type CycleDetection struct {
	g       *Graph.Graph
	visited []bool
	pre []int
	hasCycle bool
}

func NewCycleDetection(graph *Graph.Graph) *CycleDetection {
	cc := new(CycleDetection)
	cc.g = graph
	cc.hasCycle = false

	cc.visited = make([]bool, graph.V())
	cc.pre = make([]int,graph.V())

	for i:=0;i<graph.V();i++{
		if !cc.visited[i] {
			if cc.bfs(i){
				cc.hasCycle = true
				cc.bfs(i)
			}

		}
	}

	return cc
}

func (g *CycleDetection) bfs(v int) bool {
	quene := new(list.List)
	quene.PushBack(v)
	g.visited[v] = true
	g.pre[v] = v

	for quene.Front()!= nil {
		v:=quene.Remove(quene.Front()).(int)
		for _,value := range g.g.Adj(v){
			if !g.visited[value] {
				quene.PushBack(value)
				g.visited[value] = true
				g.pre[value] = v
			}else if g.pre[v] != value {
				return true
			}
		}
	}
	return false
}

func (g *CycleDetection) HasCycle() bool {
	return g.hasCycle
}

func main() {
	g := Graph.NewGraph("./05-Graph-BFS/06-Cycle-Detection-BFS/g.txt")
	cycleDetection := NewCycleDetection(g)
	fmt.Println(cycleDetection.HasCycle())

	g2 := Graph.NewGraph("./05-Graph-BFS/06-Cycle-Detection-BFS/g2.txt")
	cycleDetection2 := NewCycleDetection(g2)
	fmt.Println(cycleDetection2.HasCycle())

}