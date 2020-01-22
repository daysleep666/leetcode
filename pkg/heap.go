package pkg

// MyHeap ...
type MyHeap struct {
	List []int
}

// Len ...
func (mh *MyHeap) Len() int {
	return len(mh.List)
}

// Less ...
func (mh *MyHeap) Less(i, j int) bool {
	return mh.List[i] < mh.List[j]
}

// Swap ...
func (mh *MyHeap) Swap(i, j int) {
	mh.List[i], mh.List[j] = mh.List[j], mh.List[i]
}

// Push ...
func (mh *MyHeap) Push(x interface{}) {
	mh.List = append(mh.List, x.(int))
}

// Pop ...
func (mh *MyHeap) Pop() interface{} {
	if len(mh.List) == 0 {
		return nil
	}
	v := mh.List[len(mh.List)-1]
	mh.List = mh.List[:len(mh.List)-1]
	return v
}
