package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recu(root.Left, root.Right)
}

func recu(l1, l2 *TreeNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	if l1.Val != l2.Val {
		return false
	}
	return recu(l1.Left, l2.Right) && recu(l1.Right, l2.Left)
}

func main() {

}
