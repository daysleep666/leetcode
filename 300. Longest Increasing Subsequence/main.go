package main

import "fmt"

func lengthOfLIS(nums []int) int {
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

func main() {
	fmt.Println(lengthOfLIS([]int{-2, -1}))
	fmt.Println(lengthOfLIS([]int{-1}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
}
