package pkg

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// List ...
func (ln *ListNode) List() []int {
	var result []int
	node := ln
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

// InsertListNode ...
func InsertListNode(n *ListNode, v int) *ListNode {
	if n == nil {
		n = new(ListNode)
		n.Val = v
		return n
	}
	n.Next = InsertListNode(n.Next, v)
	return n
}

func MakeList(l []int) (head *ListNode) {
	for _, v := range l {
		head = InsertListNode(head, v)
	}
	return head
}

// GetExampleListNode 正顺序的n个
func GetExampleListNode(n int) (head *ListNode) {
	for i := 0; i < n; i++ {
		head = InsertListNode(head, i)
	}
	return head
}
