package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func reverseList(head *ListNode) *ListNode {
	var f *ListNode
	for head != nil {
		n := head.Next
		head.Next = f
		f = head
		if n == nil {
			break
		}
		head = n
	}
	return head
}

func main() {

}
