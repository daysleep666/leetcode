package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := []*TreeNode{
		root,
	}
	result := make([]int, 0, 1000)
	for len(list) != 0 {
		tmpList := make([]*TreeNode, 0, len(list)*2)
		for i := 0; i < len(list); i++ {
			result = append(result, list[i].Val)
			if list[i].Left != nil {
				tmpList = append(tmpList, list[i].Left)
			}
			if list[i].Right != nil {
				tmpList = append(tmpList, list[i].Right)
			}
		}
		list = tmpList
	}

	return result
}

func main() {

}
