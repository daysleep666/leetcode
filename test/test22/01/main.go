package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := getHigh(root)
	return b
}

func getHigh(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	a, ab := getHigh(root.Left)
	if !ab {
		return 0, false
	}
	b, bb := getHigh(root.Right)
	if !bb {
		return 0, false
	}
	if a-b > 1 || b-a > 1 {
		return 0, false
	}
	return max(a, b) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
