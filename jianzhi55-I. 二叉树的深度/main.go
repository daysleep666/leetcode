package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	if a > b {
		return a + 1
	}
	return b + 1
}

func rec(root *TreeNode, cur int, max *int) {

}

func main() {

}
