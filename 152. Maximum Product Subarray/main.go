package main

import "fmt"

func maxProduct1(nums []int) int {
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMin := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	result, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if (min*nums[i]) > (max*nums[i]) && (min*nums[i]) > (nums[i]) {
			min, max = getMin(max*nums[i], nums[i]), min*nums[i]
		} else {
			max, min = getMax(max*nums[i], nums[i]), getMin(getMin(min*nums[i], nums[i]), max*nums[i])
		}
		result = getMax(result, max)
	}
	return result
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMin := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	min := make([]int, len(nums))
	max := make([]int, len(nums))
	min[0], max[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max[i] = getMax(getMax(max[i-1]*nums[i], min[i-1]*nums[i]), nums[i])
		min[i] = getMin(getMin(max[i-1]*nums[i], min[i-1]*nums[i]), nums[i])
	}
	result := nums[0]
	for _, v := range max {
		if result < v {
			result = v
		}
	}
	return result
}

func main() {
	// fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	// fmt.Println(maxProduct([]int{-2, 0, -1}))
	// fmt.Println(maxProduct([]int{-2, 3, -4}))
	// fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
	// fmt.Println(maxProduct([]int{2, 3, -1, 1, -3, 3, 0}))
	fmt.Println(maxProduct([]int{1, 2, -1, -2, 2, 1, -2, 1}))
}
