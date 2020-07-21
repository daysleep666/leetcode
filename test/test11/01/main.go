package main

type ListNode struct {
	Val  int
	Next ListNode
}

// 双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	endA, endB := 0, 0
	for a != b && !(endA == 2 && endB == 2) {
		a = a.Next
		if a == nil {
			a = headB
			endA++
		}
		b = b.Next
		if b == nil {
			b = headA
			endB++
		}
	}
	if endA == 2 && endB == 2 {
		return nil
	}
	return a
}

func main() {

}
