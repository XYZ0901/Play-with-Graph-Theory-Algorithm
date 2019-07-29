package SingleSourcePath

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/08-PATH-Improvement/Graph"
)

type SingleSourcePath struct {
	g *Graph.Graph
	s int

	pre     []int
}

func NewSingleSourcePath(g *Graph.Graph, s int) (*SingleSourcePath) {
	gd := new(SingleSourcePath)
	g.ValidateVertex(s)
	gd.g = g
	gd.s = s
	gd.pre = make([]int, g.V())
	for i := 0; i < g.V(); i++ {
		gd.pre[i] = -1
	}

	gd.Dfs()

	return gd
}

func (g *SingleSourcePath) Dfs() {
	g.pre[g.s] = g.s
	g.dfs(g.s)
}

func (g *SingleSourcePath) dfs(v int) {
	for _, value := range g.g.Adj(v) {
		if g.pre[value] == -1 {
			g.pre[value] = v
			g.dfs(value)
		}
	}
}

func (singleSourcePath *SingleSourcePath) IsConnectedTo(t int) bool {
	singleSourcePath.g.ValidateVertex(t);
	return singleSourcePath.pre[t] != -1
}

func (singleSourcePath *SingleSourcePath) Path(t int) []int {
	res := []int{}
	if !singleSourcePath.IsConnectedTo(t) {
		return res
	}
	cur := t
	for cur != singleSourcePath.s{
		res = append(res,cur)
		cur = singleSourcePath.pre[cur]
	}
	res = append(res,cur)
	reverse(&res)
	return res
}

func reverse(res *[]int) {
	for i:=0;i<len(*res)/2;i++ {
		temp := (*res)[i]
		(*res)[i] = (*res)[len(*res)-i-1]
		(*res)[len(*res)-i-1] = temp
	}
}