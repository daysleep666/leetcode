package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

func main() {

}
