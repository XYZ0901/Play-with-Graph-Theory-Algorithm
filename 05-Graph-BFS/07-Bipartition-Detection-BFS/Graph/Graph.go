package Graph

import (
	"Play-with-Graph-Theory-Algorithm/05-Graph-BFS/07-Bipartition-Detection-BFS/tree"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 暂时只支持无向无权图
type Graph struct {
	v   int
	e   int
	adj []tree.BST
}

func (g *Graph) ValidateVertex(v int) {
	if v < 0 || v >= g.v {
		log.Fatal("vertex " + strconv.Itoa(v) + " is invalid")
	}
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) HasEdge(v, w int) bool {
	g.ValidateVertex(v)
	g.ValidateVertex(w)
	return g.adj[v].Contains(w)
}

func (g *Graph) Adj(v int) []int {
	g.ValidateVertex(v)

	return g.adj[v].Traverse()
}

func (g *Graph) Degree(v int) int {
	g.ValidateVertex(v)
	return g.adj[v].Size()
}

func NewGraph(filename string) (*Graph) {

	fileObj, _ := os.Open(filename)
	scann := bufio.NewScanner(fileObj)
	adjMatrix := new(Graph)
	firstLine := true

	var adj []tree.BST

	for scann.Scan() {
		nums := lineNums(scann)
		if firstLine {
			adjMatrix.v = nums[0]
			if adjMatrix.v < 0 {
				log.Fatal("V must be non-negative")
			}
			adjMatrix.e = nums[1]
			if adjMatrix.e < 0 {
				log.Fatal("E must be non-negative")
			}

			adj = make([]tree.BST, adjMatrix.v)

			firstLine = false
			continue
		}
		//fmt.Println(adj[6][6])
		nums = lineNums(scann)
		a := nums[0]
		adjMatrix.ValidateVertex(a)
		b := nums[1]
		adjMatrix.ValidateVertex(b)

		if a == b {
			log.Fatal("Self Loop is Detected!")
		}

		if adj[a].Contains(b) {
			log.Fatal("Parallel Edges are Detected")
		}

		adj[a].Add(b)
		adj[b].Add(a)
	}
	adjMatrix.adj = adj
	return adjMatrix
}

func (g *Graph) String() string {
	sb := ""
	sb += fmt.Sprintf("V = %d, E = %d\n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		sb += fmt.Sprintf("%d : ", i)
		sb += fmt.Sprint(g.adj[i].Traverse())
		if i < g.v-1 {
			sb += fmt.Sprintf("\n")
		}
	}
	return sb
}

func lineNums(scanner *bufio.Scanner) (nums []int) {
	line := scanner.Bytes()
	num := make([]byte, 0)
	for i := 0; i < len(line); i++ {
		if string(line[i]) == " " {
			a, _ := strconv.Atoi(string(num))
			nums = append(nums, a)
			num = make([]byte, 0)
			continue
		}
		num = append(num, line[i])
		if i == len(line)-1 {
			a, _ := strconv.Atoi(string(num))
			nums = append(nums, a)
		}
	}
	return
}
