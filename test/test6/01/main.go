package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {

}
