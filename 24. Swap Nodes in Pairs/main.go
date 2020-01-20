package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

// ListNode ...
type ListNode = pkg.ListNode

// 时间复杂度 O(n)
// 空间复杂度 O(n)
func swapPairs1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	fir := head
	sec := head.Next
	next := sec.Next
	sec.Next = fir
	fir.Next = swapPairs1(next)
	return sec
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := head.Next
	pre := &ListNode{}
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		fir := pre.Next
		sec := fir.Next
		sec.Next, fir.Next, pre.Next = fir, sec.Next, sec
		pre = fir
	}
	return newHead
}

func main() {
	n := pkg.GetExampleListNode(10)
	n = swapPairs(n)
	fmt.Println(n.List())
}

// Given a linked list, swap every two adjacent nodes and return its head.

// You may not modify the values in the list's nodes, only nodes itself may be changed.

//

// Example:

// Given 1->2->3->4, you should return the list as 2->1->4->3.
