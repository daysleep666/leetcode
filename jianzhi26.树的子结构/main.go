package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if levelOrder(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func levelOrder(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return levelOrder(a.Left, b.Left) && levelOrder(a.Right, b.Right)
}

func main() {

}
