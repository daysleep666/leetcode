package main

import (
	"fmt"
	"math"

	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

var value int

func isValidBST1(root *TreeNode) bool {
	value = int(math.MinInt64)
	return midOrder(root)
}

func midOrder(node *TreeNode) bool {
	if node == nil {
		return true
	}
	if !midOrder(node.Left) {
		return false
	}
	if node.Val <= value {
		return false
	}
	value = node.Val
	return midOrder(node.Right)
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, int(math.MinInt64), int(math.MaxInt64))
}

func isValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min {
		return false
	}
	if node.Val >= max {
		return false
	}
	return isValid(node.Left, min, node.Val) && isValid(node.Right, node.Val, max)
}

func main() {
	t := pkg.GenerateTree(2, 1, 3)
	fmt.Println(isValidBST(t))
}
