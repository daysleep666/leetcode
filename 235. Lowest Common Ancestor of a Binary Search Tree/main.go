package main

import (
	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	if root == p || root == q {
		return root
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return lowestCommonAncestor(root.Left, p, q)
}

func main() {

}
