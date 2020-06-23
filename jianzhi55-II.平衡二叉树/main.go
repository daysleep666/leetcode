package main

import "github.com/daysleep666/leetcode/pkg"

type TreeNode = pkg.TreeNode

func isBalanced(root *TreeNode) bool {
	_, b := rec(root)
	return b
}

// 返回当前节点高度和当前节点是否是平衡的
func rec(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lv, lb := rec(root.Left)
	if !lb {
		return 0, false
	}
	rv, rb := rec(root.Right)
	if !rb {
		return 0, false
	}

	if lv > rv {
		return lv + 1, (lv - rv) <= 1
	}
	return rv + 1, (rv - lv) <= 1
}

func main() {

}
