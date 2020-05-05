package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func reverseKGroup(head *ListNode, k int) (r *ListNode) {
	if head == nil || k <= 1 {
		return head
	}
	c := 1
	f := &ListNode{Next: head}
	r = f
	for head != nil && head.Next != nil {
		head = head.Next
		c++
		if c == k {
			head.Next, head = nil, head.Next
			f.Next, f = reverse(f.Next)
			f.Next = head
			c = 1
		}
	}
	return r.Next
}

func reverse(head *ListNode) (f *ListNode, l *ListNode) {
	node := head
	var n *ListNode
	for node != nil {
		n, node.Next, f = node.Next, f, node
		node = n
	}
	return f, head
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	c := 1
	f := &ListNode{Next: head}
	var l *ListNode
	for head != nil && head.Next != nil {
		head = head.Next
		c++
		if c == k {
			head.Next, head = nil, head.Next
			f.Next, l = reverse(f.Next)
			l.Next = reverseKGroup(head, k)
			break
		}
	}
	return f.Next
}

func main() {
	r := reverseKGroup(pkg.GetExampleListNode(5), 2)
	fmt.Println(r.List())
}
