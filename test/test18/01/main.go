package main

// 有小到大
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := partition(nums, start, end)
	quickSort(nums, start, mid-1)
	quickSort(nums, mid+1, end)
}

func partition(nums []int, start, end int) int {
	i, j := start, start
	for ; i < end; i++ {
		if nums[i] < nums[end] {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[end] = nums[end], nums[j]
	return j
}
