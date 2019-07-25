package AdjList

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

type AdjList struct {
	v   int
	e   int
	adj []list.List
}

func (adj *AdjList) validateVertex(v int) {
	if v < 0 || v >= adj.v {
		log.Fatal("vertex " + strconv.Itoa(v) + " is invalid")
	}
}

func (adj *AdjList) V() int {
	return adj.v
}

func (adj *AdjList) E() int {
	return adj.e
}

func (adj *AdjList) HasEdge(v, w int) bool {
	adj.validateVertex(v)
	adj.validateVertex(w)
	return ListContains(adj.adj[v],w)
}

func (adj *AdjList) Adj(v int) (res []int) {
	adj.validateVertex(v)
	for p:=adj.adj[v].Front();p!=nil;p=p.Next(){
		res = append(res, p.Value.(int))
	}
	return
}

func (adj *AdjList) Degree(v int) int {
	return len(adj.Adj(v))
}

func NewAdjList(filename string) (*AdjList) {

	fileObj, _ := os.Open(filename)
	scann := bufio.NewScanner(fileObj)
	adjMatrix := new(AdjList)
	firstLine := true

	var adj []list.List

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

			adj = make([]list.List, adjMatrix.v)

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

		if ListContains(adj[a], b) {
			log.Fatal("Parallel Edges are Detected")
		}

		adj[a].PushBack(b)
		adj[b].PushBack(a)
	}
	adjMatrix.adj = adj
	return adjMatrix
}

// golang中没有给链表contains方法 只能自己实现
func ListContains(a list.List, b interface{}) bool {
	for p := a.Front(); p != nil; p = p.Next() {
		if p.Value == b {
			return true
		}
	}
	return false
}

func (adj *AdjList) String() string {
	sb := ""
	sb += fmt.Sprintf("V = %d, E = %d\n", adj.v, adj.e)
	for i := 0; i < adj.v; i++ {
		sb += fmt.Sprintf("%d : ",i)
		for p := adj.adj[i].Front(); p != nil; p = p.Next() {
			sb += strconv.Itoa(p.Value.(int)) + " "
		}
		if i < adj.v-1 {
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
