package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 前缀和
// 如果出现了相同的数值 1 2 3 5 3 2
// 下标2==下标4 因此 下标2-下标4的和为0
func removeZeroSumSublists(head *ListNode) *ListNode {
	f := &ListNode{Next: head}
	m := make(map[int]*ListNode, 0)
	tmp := f
	sum := 0
	for tmp != nil {
		sum += tmp.Val
		m[sum] = tmp
		tmp = tmp.Next
	}

	sum = 0
	tmp = f
	for tmp != nil {
		sum += tmp.Val
		tmp.Next = m[sum].Next
		tmp = tmp.Next
	}

	return f.Next
}

// 递归?
func removeZeroSumSublists1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	f := &ListNode{Next: head}
	tmp := f
	sum := 0
	for head != nil {
		sum += head.Val
		if sum == 0 {
			tmp.Next = head.Next
		}
		head = head.Next
	}
	if f.Next != nil {
		f.Next.Next = removeZeroSumSublists(f.Next.Next)
	}
	return f.Next
}
