package main

import (
	"container/heap"
	"fmt"
	"github.com/daysleep666/leetcode/pkg"
)

// MyHeap ...
type MyHeap = pkg.MyHeap

// KthLargest ...
type KthLargest struct {
	h *MyHeap
	k int
}

// Constructor ...
func Constructor(k int, nums []int) KthLargest {
	h := &MyHeap{List: nums}
	heap.Init(h)
	for h.Len() > k {
		heap.Pop(h)
	}
	return KthLargest{
		h: h,
		k: k,
	}
}

// Add ...
func (k *KthLargest) Add(val int) int {
	heap.Push(k.h, val)
	if k.h.Len() > k.k {
		heap.Pop(k.h)
	} else if k.h.Len() == 0 {
		return 0
	}
	return k.h.List[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	fmt.Println(k.Add(9))
	fmt.Println(k.Add(4))
}

// Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

// Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.

// Example:

// int k = 3;
// int[] arr = [4,5,8,2];
// KthLargest kthLargest = new KthLargest(3, arr);
// kthLargest.add(3);   // returns 4
// kthLargest.add(5);   // returns 5
// kthLargest.add(10);  // returns 5
// kthLargest.add(9);   // returns 8
// kthLargest.add(4);   // returns 8