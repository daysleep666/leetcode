package main

import "fmt"

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i] = x + 1
	dp := make([]int, len(nums))
	dp[0] = 1
	getMin := func(i int) int {
		v := nums[i]
		max := 0
		for ; i >= 0; i-- {
			if nums[i] < v && dp[i] > max {
				max = dp[i]
			}
		}
		return max
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = getMin(i) + 1
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func lengthOfLIS1(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		v := recursive(nums[i], nums[i+1:])
		if v > max {
			max = v
		}
	}
	return max
}

func recursive(cur int, nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	if cur >= nums[0] {
		return recursive(cur, nums[1:])
	}
	return getMax(1+recursive(nums[0], nums[1:]), recursive(cur, nums[1:]))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	list := make([]int, 0, len(nums))
	getMid := func(v int) int {
		low, high := 0, len(list)-1
		mid := 0
		for low < high {
			mid = (low + high) / 2
			if list[mid] < v {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		return mid
	}
	for i := 0; i < len(nums); i++ {

	}
}

func main() {
	fmt.Println(lengthOfLIS([]int{-1}))                         //1
	fmt.Println(lengthOfLIS([]int{-2, -1}))                     //2
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) //4
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})) //6
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))          // 3
	fmt.Println(lengthOfLIS([]int{2, 2}))                       // 1
}
