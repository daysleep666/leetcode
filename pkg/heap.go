package pkg

// MyHeap ...
type MyHeap struct {
	List []int
}

// Len ...
func (mh *MyHeap) Len() int {
	return int(len(mh.List))
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
func (mh *MyHeap) Push(x int) {
	mh.List = append(mh.List, x)
}

// Pop ...
func (mh *MyHeap) Pop() int {
	if len(mh.List) == 0 {
		return 0
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
		heapify(arr, i, len(arr))
	}
}

func heapify(arr []int, i int, l int) {
	n := i
	if i*2+1 < l && arr[i] < arr[i*2+1] {
		n = i*2 + 1
	}
	if i*2+2 < l && arr[n] < arr[i*2+2] {
		n = i*2 + 2
	}
	if i == n {
		return
	}
	arr[i], arr[n] = arr[n], arr[i]
	i = n
	heapify(arr, i, l)
}

// SortByHeap 堆排序
func SortByHeap(arr []int) {
	l := len(arr)
	for l > 1 {
		arr[0], arr[l-1] = arr[l-1], arr[0]
		heapify(arr, 0, l-1)
		l--
	}
}
