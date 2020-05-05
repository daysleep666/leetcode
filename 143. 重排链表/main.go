package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var arr []*ListNode
	n := head
	for n != nil {
		t := n
		arr = append(arr, t)
		n = n.Next
	}

	n = head
	for i, j := 1, len(arr)-1; i <= j; i, j = i+1, j-1 {
		n.Next = arr[j]
		n = n.Next
		if i == j {
			break
		}
		n.Next = arr[i]
		n = n.Next
	}
	n.Next = nil
}

func main() {

}
