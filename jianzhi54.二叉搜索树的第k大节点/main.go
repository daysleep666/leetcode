package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var result int
	rec(root, &k, &result)
	return result
}

func rec(root *TreeNode, k, result *int) {
	if root == nil {
		return
	}
	if (*k) < 1 {
		return
	}
	rec(root.Right, k, result)
	(*k)--
	if (*k) == 0 {
		(*result) = root.Val
		return
	}
	rec(root.Left, k, result)
}

func main() {

}
