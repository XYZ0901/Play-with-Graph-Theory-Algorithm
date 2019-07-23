package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AdjMatrix struct {
	V   int
	E   int
	adj [][]int
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
			adjMatrix.V = nums[0]
			adjMatrix.E = nums[1]

			for i := 0; i < adjMatrix.V; i++ {
				tmp := make([]int, adjMatrix.V)
				adj = append(adj, tmp)
			}
			firstLine = false
			continue
		}
		nums = lineNums(scann)
		a := nums[0]
		b := nums[1]
		adj[a][b] = 1
		adj[b][a] = 1
	}
	adjMatrix.adj = adj
	return adjMatrix
}

func (adj *AdjMatrix) String() string {
	sb := ""
	sb += fmt.Sprintf("V = %d, E = %d\n",adj.V,adj.E)
	for i:=0;i<adj.V ; i++ {
		for j:=0;j<adj.V ; j++ {
			sb += strconv.Itoa(adj.adj[i][j])+" "
		}
		if i< adj.V-1{
			sb += fmt.Sprintf("\n")
		}
	}

	return sb
}

func main() {
	adjMatrix := NewAdjMatrix("./02-Graph-Basics/03-Adjacency-Matrix/g.txt")
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

