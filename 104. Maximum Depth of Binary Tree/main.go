package main

import (
	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

// BFS
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth int
	var list []*TreeNode
	list = append(list, root)
	for len(list) > 0 {
		depth++
		var tmp []*TreeNode
		for i := 0; i < len(list); i++ {
			n := list[i]
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}
		list = tmp
	}
	return depth
}

// DFS
func maxDepth(root *TreeNode) int {
	var max int
	dfs(root, 1, &max)
	return max
}

func dfs(node *TreeNode, depth int, max *int) {
	if node == nil {
		return
	}
	if depth > *max {
		*max = depth
	}
	dfs(node.Left, depth+1, max)
	dfs(node.Right, depth+1, max)
}

func main() {

}
