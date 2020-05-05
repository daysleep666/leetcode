package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	j, o := new(ListNode), new(ListNode)
	a, b := j, o
	trun := false

	for head != nil {
		if trun {
			// ou
			b.Next = head
			b = head
		} else {
			// ji
			a.Next = head
			a = head
		}
		trun = !trun
		head = head.Next
	}
	a.Next = o.Next
	b.Next = nil

	return j.Next
}
func main() {

}
