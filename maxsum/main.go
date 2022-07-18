package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

var res int

func newNode(data int) *Node {
	node := &Node{
		Data: data,
	}

	node.Right = nil
	node.Left = nil

	return node
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func findMaxUtil(root *Node) int {
	if root == nil {
		return 0
	}

	l := findMaxUtil(root.Left)
	r := findMaxUtil(root.Right)

	max_single := max(max(l, r)+root.Data, root.Data)
	max_top := max(max_single, l+r+root.Data)

	res = max(res, max_top)

	return max_single
}

func findMaxSum(root *Node) int {
	res = 0
	findMaxUtil(root)
	return res
}

func main() {
	root := newNode(10)
	root.Left = newNode(2)
	root.Right = newNode(10)
	root.Left.Left = newNode(20)
	root.Left.Right = newNode(1)
	root.Right.Right = newNode(-25)
	root.Right.Right.Left = newNode(3)
	root.Right.Right.Right = newNode(4)

	fmt.Printf("Max path sum is %v", findMaxSum(root))
}
