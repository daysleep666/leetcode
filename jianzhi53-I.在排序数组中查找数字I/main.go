package main

import "fmt"

// 有序数组 必然是二分
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = (low + high) / 2

		if nums[mid] == target {
			break
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if nums[mid] != target {
		return 0
	}
	c := 1
	for i := mid - 1; i >= 0; i-- {
		if nums[i] == target {
			c++
		} else {
			break
		}
	}
	for i := mid + 1; i < len(nums); i++ {
		if nums[i] == target {
			c++
		} else {
			break
		}
	}
	return c
}

func main() {
	fmt.Println(search([]int{1, 3}, 1))
}
