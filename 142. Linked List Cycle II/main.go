package main

import (
	"github.com/daysleep666/leetcode/pkg"
)

// ListNode ...
type ListNode = pkg.ListNode

func detectCycle(head *ListNode) *ListNode {
	var (
		slow = head
		fast = head
	)

	for {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			return nil
		}
		if fast == slow {
			break
		}
	}
	slow = head
	for {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
}

func main() {

}

// Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

// Note: Do not modify the linked list.
