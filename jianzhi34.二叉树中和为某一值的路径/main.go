package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	recursive(root, []int{}, 0, sum, &result)
	return result
}

func recursive(root *TreeNode, arr []int, sum, aim int, result *[][]int) {
	if root == nil {
		return
	}
	sum += root.Val
	arr = append(arr, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == aim {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			(*result) = append((*result), tmp)
		}
		return
	}
	recursive(root.Left, arr, sum, aim, result)
	recursive(root.Right, arr, sum, aim, result)
}
