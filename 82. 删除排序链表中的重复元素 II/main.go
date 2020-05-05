package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	f := &ListNode{}
	r := f
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
			continue
		}
		f.Next, f, head = head, head, head.Next
	}
	f.Next = head
	return r.Next
}

func main() {

}
