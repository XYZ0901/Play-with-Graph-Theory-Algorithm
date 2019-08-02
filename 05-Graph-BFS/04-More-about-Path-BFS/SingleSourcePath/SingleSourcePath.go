package SingleSourcePath

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/04-More-about-Path-BFS/Graph"
	"container/list"
)

type SingleSourcePath struct {
	g       *Graph.Graph
	s int
	pre  []int
	visited []bool
}

func NewSingleSourcePath(g *Graph.Graph,s int) (*SingleSourcePath) {
	gd := new(SingleSourcePath)
	gd.g = g
	gd.s = s
	for i:=0;i<g.V();i++{
		gd.pre = append(gd.pre, -1)
	}
	gd.visited = make([]bool, g.V())

	gd.bfs(s)
	return gd
}

func (g *SingleSourcePath) bfs(v int) {
	queue := new(list.List)
	queue.PushBack(v)
	g.pre[v] = v
	g.visited[v] = true
	for queue.Front() != nil {
		v := queue.Remove(queue.Front()).(int)
		for _, value := range g.g.Adj(v) {
			if !g.visited[value] {
				g.pre[value] = v
				g.visited[value] = true
				queue.PushBack(value)
			}
		}
	}
}

func (g *SingleSourcePath) IsConnectedTo (t int) bool {
	g.g.ValidateVertex(t)
	return g.visited[t]
}

func (g *SingleSourcePath) Path(t int) []int {
	res := []int{}
	if !g.IsConnectedTo(t) {
		return res
	}
	cur := t
	for cur != g.s{
		res = append(res, cur)
		cur = g.pre[cur]
	}
	res = append(res, g.s)

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