package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	{ // reverse
		var f *ListNode
		for head != nil {
			n := head.Next
			head.Next = f
			f = head
			head = n
		}
		head = f
	}
	{
		c := 1
		var f *ListNode
		for head != nil {
			if c == n {
				head = head.Next
			}
			if head == nil {
				break
			}
			c++
			next := head.Next
			head.Next = f
			f = head
			head = next
		}
		head = f
	}
	return head
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	c := 0
	node := head
	for node != nil {
		c++
		node = node.Next
	}

	n = c - n + 1

	if n == 1 {
		return head.Next
	}

	c = 0
	node = head
	for node != nil {
		c++
		if c == n-1 {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	c := 0
	fast, slow := head, head
	for fast != nil {
		if c > n {
			slow = slow.Next
		}
		c++
		fast = fast.Next
	}
	if c <= n {
		return head.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func main() {

}
