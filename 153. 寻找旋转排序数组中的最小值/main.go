package main

import "fmt"

func findMin(nums []int) int {
	// 二分
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		// fmt.Println(start, end, mid)
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}

func main() {
	fmt.Println(findMin([]int{1, 2}))                // 1
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))       // 1
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2})) // 0
}
