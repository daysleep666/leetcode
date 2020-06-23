package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	list := []*TreeNode{
		root,
	}
	positive := true
	result := make([][]int, 0, 1000)
	for len(list) != 0 {
		tmpList := make([]*TreeNode, 0, len(list)*2)
		tmpResult := make([]int, 0, len(list))
		for i := 0; i < len(list); i++ {
			if list[i].Left != nil {
				tmpList = append(tmpList, list[i].Left)
			}
			if list[i].Right != nil {
				tmpList = append(tmpList, list[i].Right)
			}
		}
		if positive {
			for i := 0; i < len(list); i++ {
				tmpResult = append(tmpResult, list[i].Val)
			}
		} else {
			for i := len(list) - 1; i >= 0; i-- {
				tmpResult = append(tmpResult, list[i].Val)
			}
		}
		list = tmpList
		result = append(result, tmpResult)
		positive = !positive
	}

	return result
}

func main() {

}
