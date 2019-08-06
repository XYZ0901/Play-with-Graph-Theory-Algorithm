package tree

type node struct {
	e     int
	left  *node
	right *node
}

func newNode(e int) *node {
	return &node{e: e, left: nil, right: nil}
}
