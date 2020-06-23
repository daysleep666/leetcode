package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	node := head
	c := 0
	for head != nil {
		if c >= k {
			node = node.Next
		}
		c++
		head = head.Next
	}
	return node
}

func main() {

}
