package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(value int) {
	s.arr = append(s.arr, value)
}

func (s *Stack) Pop() int {
	if len(s.arr) == 0 {
		return -1
	}
	l := len(s.arr)
	v := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return v
}

func (s *Stack) Count() int {
	return len(s.arr)
}

type CQueue struct {
	s1, s2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		s1: &Stack{},
		s2: &Stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	c := this.s1.Count()
	for i := 0; i < c; i++ {
		v := this.s1.Pop()
		this.s2.Push(v)
	}
	v := this.s2.Pop()
	c = this.s2.Count()
	for i := 0; i < c; i++ {
		this.s1.Push(this.s2.Pop())
	}
	return v
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	c := Constructor()
	c.AppendTail(1)
	c.AppendTail(2)
	fmt.Println(c.DeleteHead())
	fmt.Println(c.DeleteHead())
}
