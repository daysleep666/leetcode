package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		for node != nil && node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}

func main() {

}
