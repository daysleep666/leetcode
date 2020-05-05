package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	c := 1
	node := head
	for node.Next != nil {
		node = node.Next
		c++
	}
	node.Next = head

	node = head
	for i := 1; i < c-k%c; i++ {
		node = node.Next
	}
	head, node.Next = node.Next, nil
	return head
}

func main() {
	r := rotateRight(pkg.GetExampleListNode(5), 5)
	fmt.Println(r.List())
}
