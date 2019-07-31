package Path

import (
	"Play-with-Graph-Theory-Algorithm/04-Graph-DFS-Applications/11-Bipartition-Detection/Graph"
	"fmt"
)

type Path struct {
	g    *Graph.Graph
	s, t int

	visited []bool
	pre     []int
}

func NewPath(g *Graph.Graph, s, t int) (*Path) {
	this := new(Path)
	g.ValidateVertex(s)
	g.ValidateVertex(t)
	this.g = g
	this.s = s
	this.t = t

	this.visited = make([]bool,g.V())
	this.pre = make([]int, g.V())
	for i := 0; i < g.V(); i++ {
		this.pre[i] = -1
	}

	this.dfs(this.s,this.s)

	return this
}

func (path *Path) dfs(v,parent int) bool {
	path.visited[v] = true
	path.pre[v] = parent

	if path.t == v {
		return true
	}

	for _, value := range path.g.Adj(v) {
		if !path.visited[value] {
			if path.dfs(value,v){
				return true
			}
		}
	}
	return false
}

func (path *Path) IsConnectedTo() bool {
	return path.visited[path.t]
}

func (path *Path) Path() []int {
	res := []int{}

	if !path.IsConnectedTo() {
		return res
	}
	cur := path.t
	for cur != path.s {
		res = append(res, cur)
		cur = path.pre[cur]
	}
	res = append(res, cur)

	reverse(&res)
	return res
}

func main(){
	g:=Graph.NewGraph("./04-Graph-DFS-Applications/09-Cycle-Detection/g.txt")
	path := NewPath(g,0,6)
	fmt.Println("0 -> 6 : ",path.Path())

	path2 := NewPath(g,0,5)
	fmt.Println("0 -> 5 : ",path2.Path())

	path3 := NewPath(g,0,1)
	fmt.Println("0 -> 1 : ",path3.Path())
}

func reverse(res *[]int) {
	for i := 0; i < len(*res)/2; i++ {
		temp := (*res)[i]
		(*res)[i] = (*res)[len(*res)-i-1]
		(*res)[len(*res)-i-1] = temp
	}
}
