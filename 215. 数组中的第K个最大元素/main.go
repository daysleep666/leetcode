package main

// 堆
func findKthLargest(nums []int, k int) int {
	// 构造小顶堆
	for i := k / 2; i >= 0; i-- {
		heapify(nums, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[i], nums[0] = nums[0], nums[i]
		}
		heapify(nums, 0, k)
	}
	return nums[0]
}

func heapify(arr []int, start, end int) {
	cur := start

	if start*2+1 < end && arr[cur] > arr[start*2+1] {
		cur = start*2 + 1
	}
	if start*2+2 < end && arr[cur] > arr[start*2+2] {
		cur = start*2 + 2
	}

	if cur == start {
		return
	}
	arr[cur], arr[start] = arr[start], arr[cur]
	heapify(arr, cur, end)
}

func main() {

}
