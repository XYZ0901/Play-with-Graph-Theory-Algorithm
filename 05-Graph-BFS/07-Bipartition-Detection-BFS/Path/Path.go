package Path

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/07-Bipartition-Detection-BFS/Graph"
	"container/list"
	"fmt"
)

type Path struct {
	g       *Graph.Graph
	s,t int
	pre  []int
	visited []bool
}

func NewPath(g *Graph.Graph,s,t int) (*Path) {
	g.ValidateVertex(s)
	g.ValidateVertex(t)

	gd := new(Path)
	gd.g = g
	gd.s,gd.t = s,t
	for i:=0;i<g.V();i++{
		gd.pre = append(gd.pre, -1)
	}
	gd.visited = make([]bool, g.V())

	gd.bfs()
	return gd
}

func (g *Path) bfs() {
	queue := new(list.List)
	queue.PushBack(g.s)
	g.pre[g.s] = g.s
	g.visited[g.s] = true
	if g.s == g.t {
		return
	}

	for queue.Front() != nil {
		v := queue.Remove(queue.Front()).(int)
		for _, value := range g.g.Adj(v) {
			if !g.visited[value] {
				g.pre[value] = v
				g.visited[value] = true
				queue.PushBack(value)
				if value == g.t {
					return
				}
			}
		}
	}
}

func (g *Path) IsConnectedTo (t int) bool {
	return g.visited[t]
}

func (g *Path) Path() []int {
	res := []int{}
	if !g.IsConnectedTo(g.t) {
		return res
	}
	cur := g.t
	for cur != g.s{
		res = append(res, cur)
		cur = g.pre[cur]
	}
	res = append(res, g.s)

	reverse(&res)
	return res
}

func main() {
	g := Graph.NewGraph("./05-Graph-BFS/02-BFS/g.txt")
	graphBFS := NewPath(g,0,6)
	fmt.Println("0 -> 6 : ",graphBFS.Path())

	graphBFS2 := NewPath(g,0,5)
	fmt.Println("0 -> 5 : ",graphBFS2.Path())

	graphBFS3 := NewPath(g,0,1)
	fmt.Println("0 -> 1 : ",graphBFS3.Path())
}

func reverse(res *[]int) {
	for i:=0;i<len(*res)/2;i++ {
		temp := (*res)[i]
		(*res)[i] = (*res)[len(*res)-i-1]
		(*res)[len(*res)-i-1] = temp
	}
}
