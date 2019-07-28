package main

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/07-All-Pairs-Path/Graph"
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/07-All-Pairs-Path/SingleSourcePath"
	"fmt"
)

type AllPairsPath struct {
	G     *Graph.Graph
	paths []SingleSourcePath.SingleSourcePath
}

func NewAllPairsPath(G *Graph.Graph) *AllPairsPath {
	allPairsPath := new(AllPairsPath)
	allPairsPath.G = G
	allPairsPath.paths = make([]SingleSourcePath.SingleSourcePath,G.V())

	for i:=0;i<G.V() ; i++ {
		allPairsPath.paths[i] = *SingleSourcePath.NewSingleSourcePath(G,i)
	}

	return allPairsPath
}

func(allPairsPath*AllPairsPath)IsConnectedTo(s,t int) bool {
	allPairsPath.G.ValidateVertex(s)
	return allPairsPath.paths[s].IsConnectedTo(t)
}

func(allPairsPath*AllPairsPath)Path(s,t int) []int {
	allPairsPath.G.ValidateVertex(s)
	return allPairsPath.paths[s].Path(t)
}

func main() {
	g := Graph.NewGraph("./04-Graph-DFS-Applications/07-All-Pairs-Path/g.txt")
	all := NewAllPairsPath(g)
	fmt.Println(all.Path(0,6))
	fmt.Println(all.Path(1,6))
	fmt.Println(all.Path(4,6))
	fmt.Println(all.Path(5,6))

}
