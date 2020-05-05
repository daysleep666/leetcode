package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	l, b := &ListNode{}, &ListNode{}
	ls, bs := l, b
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = head
		} else {
			b.Next = head
			b = head
		}
		head = head.Next
	}
	b.Next = nil
	l.Next = bs.Next
	return ls.Next
}

func main() {
	r := partition(pkg.MakeList([]int{1, 4, 3, 2, 5, 2}), 3)
	fmt.Println(r.List())
}
