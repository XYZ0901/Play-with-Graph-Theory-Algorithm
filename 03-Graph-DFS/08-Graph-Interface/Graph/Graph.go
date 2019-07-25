package Graph

type Graph interface {
	V() int
	E() int
	HasEdge(v,m int) bool
	Adj(v int) []int
	Degree(v int) int
}
