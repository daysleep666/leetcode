package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的中序遍历是有序数组
func findMode1(root *TreeNode) []int {
	result := make([]int, 0)
	cur, max := 0, 0
	recursive(root, &result, &cur, &max)
	{
		if (cur) > (max) {
			(result) = (result)[len((result))-1:]
			(max) = (cur)
		} else if (cur) < (max) {
			(result) = (result)[:len((result))-1]
		}
	}
	return result
}

func recursive(root *TreeNode, result *[]int, cur, max *int) {
	if root == nil {
		return
	}
	recursive(root.Left, result, cur, max)
	{
		last := (*result)[len((*result))-1]
		if root.Val == last {
			(*cur)++
		}

		if len(*result) == 0 {
			(*result) = append((*result), root.Val)
			(*max) = 1
			(*cur) = 1
		} else {
			last := (*result)[len((*result))-1]
			if root.Val == last {
				(*cur)++
			}
		}

		if (*cur) > (*max) {
			(*result) = (*result)[len((*result))-1:]
			(*max) = (*cur)
		} else if (*cur) < (*max) {
			(*result) = (*result)[:len((*result))-1]
		}
		(*result) = append((*result), root.Val)
		(*cur) = 1

	}
	recursive(root.Right, result, cur, max)
}

// 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
