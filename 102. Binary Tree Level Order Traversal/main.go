package main

import (
	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

func levelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}
	var list []*TreeNode
	list = append(list, root)
	for len(list) > 0 {
		var tmp []*TreeNode
		add := func(ns ...*TreeNode) {
			for _, n := range ns {
				if n != nil {
					tmp = append(tmp, n)
				}
			}
		}
		r := make([]int, 0, len(list))
		for i := 0; i < len(list); i++ {
			n := list[i]
			r = append(r, n.Val)
			add(n.Left, n.Right)
		}
		list = tmp
		result = append(result, r)
	}
	return
}

var result [][]int

func add(level, v int) {
	for level >= len(result) {
		result = append(result, []int{})
	}
	result[level] = append(result[level], v)
}

// DFS 实现
func levelOrder1(root *TreeNode) [][]int {
	result = make([][]int, 0)
	dfsLevelOrder(root, 0)
	return result
}

func dfsLevelOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	add(level, root.Val)
	dfsLevelOrder(root.Left, level+1)
	dfsLevelOrder(root.Right, level+1)
}
