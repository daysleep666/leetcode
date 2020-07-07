package main

type ListNode struct {
	Val  int
	Next ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	c := 0
	for fast != nil {
		if c > n {
			slow = slow.Next
		}
		fast = fast.Next
		c++
	}
	if c == n {
		return head.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	} else {
		slow.Next = nil
	}
	return head
}

func main() {

}
