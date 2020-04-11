package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	tmp := result
	value := 0

	add := func(sum int) {
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		sum = sum + value
		tmp.Val = sum % 10
		value = 0
		if sum >= 10 {
			value = 1
		}
	}

	for l1 != nil && l2 != nil {
		add(l1.Val + l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		add(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		add(l2.Val)
		l2 = l2.Next
	}
	if value == 1 {
		add(0)
	}
	return (result.Next)
}

func main() {
	{
		l1 := new(ListNode)
		l1 = pkg.InsertListNode(l1, 0)
		l1 = pkg.InsertListNode(l1, 8)

		l2 := new(ListNode)
		l2 = pkg.InsertListNode(l2, 6)
		l2 = pkg.InsertListNode(l2, 7)
		r := addTwoNumbers(l1.Next, l2.Next)
		fmt.Println(r.List())
	}

	{
		l1 := new(ListNode)
		l1 = pkg.InsertListNode(l1, 1)
		l1 = pkg.InsertListNode(l1, 8)

		l2 := new(ListNode)
		l2 = pkg.InsertListNode(l2, 0)
		r := addTwoNumbers(l1.Next, l2.Next)
		fmt.Println(r.List())
	}
}
