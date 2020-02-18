package main

import (
	"fmt"
	"github.com/daysleep666/leetcode/pkg"
)

// TreeNode ...
type TreeNode = pkg.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	}
	return r
}

func test(a *int) {
	fmt.Printf("test:%p\n", a)
	b := 1
	*a = b
	fmt.Println(a, *a)
}

func main() {
	var a int
	fmt.Printf("main:%p\n", &a)
	test(&a)
	fmt.Println(a)
}
