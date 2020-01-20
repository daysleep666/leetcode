package main

import (
	"fmt"
	"github.com/daysleep666/leetcode/pkg"
)

// ListNode ...
type ListNode = pkg.ListNode

// 时间复杂度 O(n)
// 空间复杂度 O(1)
// iteratively
func reverseList1(head *ListNode) *ListNode {
	var front *ListNode
	for head != nil {
		next := head.Next
		head.Next = front
		front = head
		head = next
	}
	return front
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)
// recursively
func reverseList(head *ListNode) *ListNode {
	return recursively(head, nil)
}

func recursively(curNode, frontNode *ListNode) *ListNode {
	if curNode == nil {
		return nil
	}
	next := curNode.Next
	curNode.Next = frontNode
	if next == nil {
		return curNode
	}
	return recursively(next, curNode)
}

func main() {
	n := pkg.GetExampleListNode(10)
	n = reverseList(n)
	fmt.Println(n.List())
}

// Reverse a singly linked list.

// Example:

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
// Follow up:

// A linked list can be reversed either iteratively or recursively. Could you implement both?
