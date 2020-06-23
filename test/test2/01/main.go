package main

import "github.com/daysleep666/leetcode/pkg"

type ListNode = pkg.ListNode

// 考点: 快慢指针, 反转单链表, 合并链表
func reorderList(head *ListNode) {
	if head == nil {
		return nil
	}
	// 通过快慢指针找到中心节点
	// 1 2 3 4 mid = 2
	// 1 2 3 4 5 mid = 3
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}

	node := slow.Next
	slow.Next = nil
	// 将后半部分节点反转
	var last *ListNode
	for node != nil {
		n := node.Next
		node.Next = last
		last = node
		node = n
	}

	// 合并链表
	head.Next = combine(last, head.Next)
}

func combine(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	l3 := l
	for l1 != nil && l2 != nil {
		l3.Next = l1
		l1 = l1.Next
		l3 = l3.Next

		l3.Next = l2
		l2 = l2.Next
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return l.Next
}

func main() {

}

// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例 1:

// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
// 示例 2:

// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
