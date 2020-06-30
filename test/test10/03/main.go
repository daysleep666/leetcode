package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	front := &ListNode{}
	result := front
	front.Next = head
	node := head

	for node != nil {
		delSelf := false
		for node != nil && node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
			delSelf = true
		}
		if delSelf {
			front.Next = node.Next
		} else {
			front = node
		}
		node = node.Next

	}

	return result.Next
}

func main() {

}
