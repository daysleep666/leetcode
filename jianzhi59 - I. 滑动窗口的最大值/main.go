package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	var info struct {
		max   int // 最大值
		index int // 最大值下标
		list  []int
	}
	info.list = make([]int, 0, len(nums))
	start := 0
	end := k - 1
	cur := 0
	for cur < len(nums) {
		for ; cur <= end; cur++ {
			if start > info.index {
				cur = start
				info.max = nums[cur]
				info.index = start
				continue
			}
			if nums[cur] > info.max {
				info.max = nums[cur]
				info.index = cur
			}
		}
		start++
		end++
		info.list = append(info.list, info.max)
	}
	return info.list
}

func main() {
	// // 3,3,5,5,6,7
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
