package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历是顺序的
func minDiffInBST(root *TreeNode) int {
	min := math.MaxInt64
	inOrder(root, new(*TreeNode), &min)
	return min
}

func inOrder(root *TreeNode, preNode **TreeNode, min *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, preNode, min)

	if (*preNode) != nil {
		tmp := root.Val - (*preNode).Val
		if (*min) > tmp {
			(*min) = tmp
		}
	}

	(*preNode) = &root
	inOrder(root.Right, preNode, min)
}
