package pkg

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GenerateTree ...
func GenerateTree(list ...int) *TreeNode {
	var head *TreeNode
	for _, v := range list {
		head = InsertNode(head, v)
	}
	return head
}

// InsertNode ...
func InsertNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		node.Left = InsertNode(node.Left, val)
	} else {
		node.Right = InsertNode(node.Right, val)
	}
	return node
}

// Print ...
func (n *TreeNode) Print() {
	list := make([]*TreeNode, 0)
	list = append(list, n)
	for len(list) > 0 {
		tmpList := make([]*TreeNode, 0)
		for _, v := range list {
			if v == nil {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d ", v.Val)
				tmpList = append(tmpList, v.Left, v.Right)
			}
		}
		fmt.Println()
		list = tmpList
	}
}
