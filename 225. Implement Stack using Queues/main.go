package main

import "fmt"

// MyStack ...
type MyStack struct {
	list []int
}

// Constructor /** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

// Push /** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	s.list = append(s.list, x)
}

// Pop /** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	if len(s.list) == 0 {
		return 0
	}
	v := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return v
}

// Top /** Get the top element. */
func (s *MyStack) Top() int {
	if len(s.list) == 0 {
		return 0
	}
	return s.list[len(s.list)-1]
}

// Empty /** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	return len(s.list) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	p2 := obj.Pop()
	p3 := obj.Top()
	p4 := obj.Empty()
	fmt.Println(p2, p3, p4)
}
