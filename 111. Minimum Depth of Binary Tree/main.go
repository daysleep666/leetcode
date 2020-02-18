package main

import (
	"math"

	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

func minDepth1(root *TreeNode) int {
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
			if n.Left == nil && n.Right == nil {
				return depth
			}
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := int(math.MaxInt64)
	dfs(root, 1, &min)
	return min
}

func dfs(node *TreeNode, depth int, min *int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil && depth < *min {
		*min = depth
		return
	}
	dfs(node.Left, depth+1, min)
	dfs(node.Right, depth+1, min)
}
