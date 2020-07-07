package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	i := len(nums) / 2
	head := &TreeNode{Val: nums[i]}
	head.Left = sortedArrayToBST(nums[:i])
	head.Right = sortedArrayToBST(nums[i+1:])
	return head
}
