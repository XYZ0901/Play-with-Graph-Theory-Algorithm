package main

import (
	"bufio"
	"fmt"
	"Play-with-Graph-Theory-Algorithm/02-Graph-Basics/09-Graph-Representation-Comparation/tree"
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

func (adj *Graph) validateVertex(v int) {
	if v < 0 || v >= adj.v {
		log.Fatal("vertex " + strconv.Itoa(v) + " is invalid")
	}
}

func (adj *Graph) V() int {
	return adj.v
}

func (adj *Graph) E() int {
	return adj.e
}

func (adj *Graph) HasEdge(v, w int) bool {
	adj.validateVertex(v)
	adj.validateVertex(w)
	return adj.adj[v].Contains(w)
}

func (adj *Graph) Adj(v int) interface{} {
	adj.validateVertex(v)

	return adj.adj[v]
}

func (adj *Graph) Degree(v int) int {
	adj.validateVertex(v)
	return adj.adj[v].Size()
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
		adjMatrix.validateVertex(a)
		b := nums[1]
		adjMatrix.validateVertex(b)

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

func main() {
	g := NewGraph("./02-Graph-Basics/09-Graph-Representation-Comparation/g.txt")
	fmt.Println(g)
}

// 简单封装 读取每一行以" "为分割的数字
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
