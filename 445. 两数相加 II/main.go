package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var a1, a2 []int
	n := l1
	for n != nil {
		a1 = append(a1, n.Val)
		n = n.Next
	}
	n = l2
	for n != nil {
		a2 = append(a2, n.Val)
		n = n.Next
	}
	length := len(a1)
	if len(a2) > length {
		length = len(a2)
	}
	var head *ListNode
	var jin int
	for i, j := len(a1)-1, len(a2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		a, b := 0, 0
		if i >= 0 {
			a = a1[i]
		}
		if j >= 0 {
			b = a2[j]
		}
		sum := a + b + jin
		n := &ListNode{
			Val: sum % 10,
		}
		jin = sum / 10
		n.Next = head
		head = n
	}
	if jin == 1 {
		n := &ListNode{
			Val: 1,
		}
		n.Next = head
		head = n
	}
	return head
}
func main() {

}
