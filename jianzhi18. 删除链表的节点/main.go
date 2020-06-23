package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	var front *ListNode
	node := head
	for node != nil {
		if node.Val == val {
			front.Next = node.Next
			return head
		}
		front = node
		node = node.Next
	}
	return head
}

func main() {

}
