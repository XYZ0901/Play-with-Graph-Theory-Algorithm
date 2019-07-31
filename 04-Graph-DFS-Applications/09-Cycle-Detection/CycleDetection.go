package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/09-Cycle-Detection/Graph"
	"fmt"
)

type CycleDetection struct {
	g       *Graph.Graph
	visited []bool
	hasCycle bool
}

func NewCycleDetection(g *Graph.Graph) (*CycleDetection) {
	gd := new(CycleDetection)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.hasCycle = false

	for i := 0; i < gd.g.V(); i++ {
		if !gd.visited[i] {
			if gd.dfs(i,i){
				gd.hasCycle = true
				break
			}
		}
	}
	return gd
}

func (g *CycleDetection) dfs(v,parent int) bool {
	g.visited[v] = true

	for _, value := range g.g.Adj(v) {
		if !g.visited[value] {
			if g.dfs(value,v) {
				return true
			}
		}else if value!=parent {
			return true
		}
	}
	return false
}

func (g *CycleDetection) HasCycle() bool {
	return g.hasCycle
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/09-Cycle-Detection/g.txt")
	cycleDetection := NewCycleDetection(g)
	fmt.Println(cycleDetection.hasCycle)

	g2 := Graph.NewGraph("./04-Graph-DFS-Applications/09-Cycle-Detection/g2.txt")
	cycleDetection2 := NewCycleDetection(g2)
	fmt.Println(cycleDetection2.hasCycle)
}
