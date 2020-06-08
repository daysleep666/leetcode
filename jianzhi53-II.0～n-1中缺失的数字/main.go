package main

import "fmt"

func missingNumber(nums []int) int {
	low, high, mid := 0, len(nums)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 2, 3}))
	fmt.Println(missingNumber([]int{0}))
	fmt.Println(missingNumber([]int{1, 2}))
	fmt.Println(missingNumber([]int{0, 2}))
	fmt.Println(missingNumber([]int{0, 1, 2, 3, 5}))
}
