package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func sortList(head *ListNode) *ListNode {
	// 归并
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val > head.Next.Val {
			head.Next.Next = head
			n := head.Next
			head.Next = nil
			return n
		}
		return head
	}
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	n := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(n)

	r := &ListNode{}
	result := r
	for l1 != nil || l2 != nil {
		if l1 == nil || (l2 != nil && l1.Val > l2.Val) {
			r.Next = l2
			l2 = l2.Next
			r = r.Next
			continue
		}
		r.Next = l1
		l1 = l1.Next
		r = r.Next
	}

	return result.Next
}

func main() {
	l := pkg.GetExampleListNode(5)
	l = sortList(l)
	fmt.Println(l.List())
}

