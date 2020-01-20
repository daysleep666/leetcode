package main

import (
	"github.com/daysleep666/leetcode/pkg"
)

// ListNode ...
type ListNode = pkg.ListNode

func hasCycle(head *ListNode) bool {
	var (
		slow = head
		fast = head
	)

	for {
		if slow != nil {
			slow = slow.Next
		}
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}
		if slow == fast {
			return true
		}
	}
}

func main() {

}

// Given a linked list, determine if it has a cycle in it.

// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
