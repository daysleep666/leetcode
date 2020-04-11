package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	head := l
	for l1 != nil || l2 != nil {
		if l2 == nil || l1 != nil && l1.Val < l2.Val {
			l.Next, l1 = l1, l1.Next
		} else {
			l.Next, l2 = l2, l2.Next
		}
		l = l.Next
	}
	return head.Next
}

func main() {

}
