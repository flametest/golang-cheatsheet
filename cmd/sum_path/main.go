package main

import (
	"fmt"
	"math"
)

type Node struct {
	right *Node
	left *Node
	value int
}
func dfs(root *Node) int {
	if root == nil {
		return 0
	}
	value := root.value
	leftSum := 0
	rightSum := 0
	if root.left != nil {
		leftSum = dfs(root.left)
	}
	if root.right != nil {
		rightSum = dfs(root.right)
	}
	return value + int(math.Max(float64(leftSum), float64(rightSum)))
}

func main() {
	root := Node{value: 2}
	root.left = &Node{value: 3}
	root.right = &Node{value: 4}
	root.right.left = &Node{value: 5}
	root.right.right = &Node{value: 7}
	root.right.left.left = &Node{value: 6}

	fmt.Println(dfs(&root))
}
