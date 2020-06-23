package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	ll := l
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}

	return ll.Next
}

func main() {

}
