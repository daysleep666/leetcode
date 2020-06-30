package main

type CQueue struct {
	arr1 []int
	arr2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.arr1 = append(this.arr1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.arr2) == 0 {
		for i := len(this.arr1) - 1; i >= 0; i-- {
			this.arr2 = append(this.arr2, this.arr1[i])
		}
		this.arr1 = this.arr1[:0]
	}
	if len(this.arr2) == 0 {
		return -1
	}

	v := this.arr2[len(this.arr2)-1]
	this.arr2 = this.arr2[:len(this.arr2)-1]
	return v
}

func main() {

}