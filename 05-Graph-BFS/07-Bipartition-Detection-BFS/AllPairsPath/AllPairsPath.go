package AllPairsPath

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/07-Bipartition-Detection-BFS/Graph"
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/07-Bipartition-Detection-BFS/SingleSourcePath"
)

type AllPairsPath struct {
	g     *Graph.Graph
	paths []SingleSourcePath.SingleSourcePath
}

func NewAllPairsPath(g *Graph.Graph) *AllPairsPath {
	gd := new(AllPairsPath)
	gd.g = g
	gd.paths = make([]SingleSourcePath.SingleSourcePath, g.V())

	for i := 0; i < g.V(); i++ {
		gd.paths[i] = *SingleSourcePath.NewSingleSourcePath(g, i)
	}

	return gd
}

func (g *AllPairsPath) IsConnectedTo(s, t int) bool {
	g.g.ValidateVertex(s)
	return g.paths[s].IsConnectedTo(t)
}

func (g *AllPairsPath) path(s, t int) []int {
	g.g.ValidateVertex(s)
	return g.paths[s].Path(t)
}
