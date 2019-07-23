package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AdjMatrix struct {
	v   int
	e   int
	adj [][]int
}

func (adj *AdjMatrix) validateVertex(v int) {
	if v < 0 || v >= adj.v {
		log.Fatal("vertex " + strconv.Itoa(v) + " is invalid")
	}
}

func (adj *AdjMatrix) V() int {
	return adj.v
}

func (adj *AdjMatrix) E() int {
	return adj.e
}

func (adj *AdjMatrix) HasEdge(v, w int) bool {
	adj.validateVertex(v)
	adj.validateVertex(w)
	return adj.adj[v][w] == 1
}

func (adj *AdjMatrix)Adj(v int) (res []int){
	adj.validateVertex(v)

	for i:=0;i<adj.v ;i++  {
		if adj.adj[adj.v][i]==1 {
			res = append(res, i)
		}
	}

	return res
}

func (adj *AdjMatrix)Degree(v int) int{
	return len(adj.Adj(v))
}

func NewAdjMatrix(filename string) (*AdjMatrix) {
	fileObj, _ := os.Open(filename)
	scann := bufio.NewScanner(fileObj)
	adjMatrix := new(AdjMatrix)
	firstLine := true
	var adj [][]int

	for scann.Scan() {
		nums := lineNums(scann)
		if firstLine {
			adjMatrix.v = nums[0]
			if adjMatrix.v < 0 {
				//panic("V must be non-negative")
				log.Fatal("V must be non-negative")
			}
			adjMatrix.e = nums[1]

			for i := 0; i < adjMatrix.v; i++ {
				tmp := make([]int, adjMatrix.e)
				adj = append(adj, tmp)
			}
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

		if adj[a][b] == 1 {
			log.Fatal("Parallel Edges are Detected")
		}

		adj[a][b] = 1
		adj[b][a] = 1
	}
	adjMatrix.adj = adj
	return adjMatrix
}

func (adj *AdjMatrix) String() string {
	sb := ""
	sb += fmt.Sprintf("V = %d, E = %d\n", adj.v, adj.e)
	for i := 0; i < adj.v; i++ {
		for j := 0; j < adj.v; j++ {
			sb += strconv.Itoa(adj.adj[i][j]) + " "
		}
		if i < adj.v-1 {
			sb += fmt.Sprintf("\n")
		}
	}

	return sb
}

func main() {

	adjMatrix := NewAdjMatrix("./02-Graph-Basics/08-Adjacency-Set/g.txt")
	fmt.Println(adjMatrix)
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
