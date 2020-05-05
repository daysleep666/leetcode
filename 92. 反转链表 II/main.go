package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	node, f := head, &ListNode{Next: head}
	c := 1
	r := f
	for node != nil {
		if c < m {
			f = node
		}
		if c == n {
			node.Next, node = nil, node.Next
			f.Next, f = reverse(f.Next)
			f.Next = node
		}
		c++
		if node == nil {
			break
		}
		node = node.Next
	}

	return r.Next
}

func reverse(head *ListNode) (f *ListNode, l *ListNode) {
	node := head
	var n *ListNode
	for node != nil {
		n, node.Next, f = node.Next, f, node
		node = n
	}
	return f, head
}

func main() {

}
