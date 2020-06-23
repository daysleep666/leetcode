package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{}, 100)
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}

	for headB != nil {
		_, ok := m[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 双指针
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}
		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

func main() {

}
