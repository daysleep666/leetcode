package main

import (
	"fmt"

	"github.com/daysleep666/leetcode/pkg"
)

type ListNode = pkg.ListNode

func mergeKLists1(lists []*ListNode) *ListNode {
	// 过滤空指针
	tmp := make([]*ListNode, 0, len(lists))
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			tmp = append(tmp, lists[i])
		}
	}
	lists = tmp
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	head := new(ListNode)
	l := head
	get := func() *ListNode {
		min := 0
		for i := 1; i < len(lists); i++ {
			if lists[min].Val > lists[i].Val {
				lists[min] = lists[i]
				min = i
			}
		}
		lists[min] = lists[min].Next
		if lists[min] == nil {
			lists = append(lists[:min], lists[min+1:]...)
		}
		return lists[min]
	}
	for len(lists) <= 1 {
		l.Next = get()
		l = l.Next
	}
	if len(lists) == 1 {
		l.Next = lists[0]
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	return mergeKLists([]*ListNode{mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:])})
}

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
	{
		l1 := new(ListNode)
		l1 = pkg.InsertListNode(l1, 1)
		l1 = pkg.InsertListNode(l1, 4)

		l2 := new(ListNode)
		l2 = pkg.InsertListNode(l2, 1)
		l2 = pkg.InsertListNode(l2, 3)
		r := mergeKLists([]*ListNode{l1, l2})
		fmt.Println(r.List())
	}
}
