package main

type MinStack struct {
	arr    []int
	minarr []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.minarr) == 0 || this.minarr[len(this.minarr)-1] > x {
		this.minarr = append(this.minarr, x)
	}
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		return
	}
	v := this.arr[:len(this.arr)-1]
	if v == this.minarr[len(this.minarr)-1] {
		this.minarr = this.minarr[:len(this.minarr)-1]
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		return 0
	}
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) Min() int {
	if len(this.minarr) == 0 {
		return 0
	}
	return this.minarr[len(this.minarr)-1]
}

func main() {

}
