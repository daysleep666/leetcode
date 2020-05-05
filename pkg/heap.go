package pkg

// MyHeap ...
type MyHeap struct {
	List []interface{}
}

// Len ...
func (mh *MyHeap) Len() int {
	return int(len(mh.List))
}

// Less ...
func (mh *MyHeap) Less(i, j int) bool {
	switch mh.List[i].(type) {
	case int:
		return mh.List[i].(int) < mh.List[j].(int)
	}
	return false
}

// Swap ...
func (mh *MyHeap) Swap(i, j int) {
	mh.List[i], mh.List[j] = mh.List[j], mh.List[i]
}

// Push ...
func (mh *MyHeap) Push(x interface{}) {
	mh.List = append(mh.List, x.(interface{}))
}

// Pop ...
func (mh *MyHeap) Pop() interface{} {
	if len(mh.List) == 0 {
		return nil
	}
	v := mh.List[0]
	mh.List[0] = mh.List[mh.Len()-1]
	mh.List = mh.List[:len(mh.List)-1]
	// mh.Build()
	return v
}

// Heapify 重建最小堆
func Heapify(arr []int) {
	// (index-1)/2
	for i := len(arr) / 2; i >= 0; i-- {
		if i*2+1 < len(arr) && arr[i] > arr[i*2+1] {
			arr[i], arr[i*2+1] = arr[i*2+1], arr[i]
		}
		if i*2+2 < len(arr) && arr[i] > arr[i*2+2] {
			arr[i], arr[i*2+2] = arr[i*2+2], arr[i]
		}
	}
}
