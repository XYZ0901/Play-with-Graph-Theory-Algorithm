package main

import (
	"Play-with-Graph-Theory-Algorithm/03-Graph-DFS/09-Graph-DFS-without-Recursion/Graph"
	"fmt"
)

type GraphDFSnr struct {
	g       *Graph.Graph
	visited []bool
	pre     []int
}

func NewGraphDFSnr(g *Graph.Graph) (*GraphDFSnr) {
	gd := new(GraphDFSnr)
	gd.g = g
	gd.visited = make([]bool, g.V())
	gd.Dfs()
	return gd
}

func (g *GraphDFSnr) Dfs() {
	for i := 0; i < g.g.V(); i++ {
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

func (g *GraphDFSnr) dfs(v int) {
	stack := Stack{}
	stack.Push(v)
	g.visited[v] = true
	for !stack.IsEmpty(){
		a := stack.Pop().(int)
		g.pre = append(g.pre, a)
		for _,value := range g.g.Adj(a){
			if !g.visited[value]{
				stack.Push(value)
				g.visited[value] = true
			}
		}
	}
}

func (g *GraphDFSnr) Pre() []int {
	return g.pre
}

func main() {
	g := Graph.NewGraph("./03-Graph-DFS/09-Graph-DFS-without-Recursion/g.txt")
	graphDFS := NewGraphDFSnr(g)
	fmt.Println(graphDFS.Pre())
}


// 自己实现栈
type Stack []interface {}

func (stack Stack) IsEmpty() bool  {
	return len(stack) == 0
}

func (stack *Stack) Push(value interface{})  {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (interface{})  {
	value := (*stack)[len(*stack) - 1]
	*stack = (*stack)[:len(*stack) - 1]
	return value
}
