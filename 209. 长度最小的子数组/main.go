package main

import "fmt"

// 滑动窗口 双指针
// 时间复杂度O(n) start和end各移动n次
func minSubArrayLen(s int, nums []int) int {
	start, end := 0, 0
	sum := 0
	min := len(nums) + 1
	for start <= end && end < len(nums) {
		sum += nums[end]
		for sum >= s && start <= end && end < len(nums) {
			if (end-start)+1 < min {
				min = (end - start) + 1
			}
			sum = sum - nums[start]
			start++
		}
		end++
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(3, []int{1, 1}))
}
