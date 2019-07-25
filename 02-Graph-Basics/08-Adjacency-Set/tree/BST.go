package tree

import (
	"fmt"
)

type BST struct {
	root *node
	size int
}

func NewBST() *BST {
	return &BST{root: nil, size: 0}
}

func (bst *BST) Size() int {
	return bst.size
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BST) Add(e int) {
	bst.root = bst.add(bst.root, e)
}

func (bst *BST) add(node *node, e int) *node {
	if node == nil {
		bst.size++
		return newNode(e)
	}

	if e < node.e {
		node.left = bst.add(node.left, e)
	} else if e > node.e {
		node.right = bst.add(node.right, e)
	}
	return node
}

func (bst *BST) Contains(e int) bool {
	return contains(bst.root, e)
}

func (bst *BST) Traverse() []int {
	res := []int{}
	traverse(bst.root,&res)
	return res
}

func (bst *BST) String() string {
	res := ""
	generateBSTString(bst.root,0, &res)
	return res
}

func generateBSTString(node *node, depth int, res *string) {
	if node == nil{
		*res+=generateDepthString(depth)+fmt.Sprintf("null\n")
		return
	}

	*res += generateDepthString(depth)+fmt.Sprintf("%d\n",node.e)
	generateBSTString(node.left,depth+1,res)
	generateBSTString(node.right,depth+1,res)
	return
}

func generateDepthString(depth int) string {
	res := ""
	for i:=0;i<depth;i++{
		res+="--"
	}
	return res
}



func contains(node *node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.e {
		return true
	} else if e < node.e {
		return contains(node.left, e)
	} else if e > node.e {
		return contains(node.right, e)
	}
	return false
}

func traverse(node *node,res *[]int) {
	if node!=nil {
		if node.left!=nil {
			traverse(node.left,res)
		}
		*res = append(*res, node.e)
		if node.right!=nil {
			traverse(node.right,res)
		}
	}
}

