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

func lengthOfLIS2(nums []int) int {
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
	if len(nums) <= 1 {
		return len(nums)
	}
	list := make([]int, 0, len(nums))
	getMid := func(v int) int {
		low, high := 0, len(list)
		mid := 0
		for low < high {
			mid = (low + high) / 2
			if list[mid] == v {
				return mid
			} else if list[mid] < v {
				low = mid + 1
			} else {
				high = mid
			}
		}
		return low
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if len(list) == 0 {
			list = append(list, nums[i])
			continue
		}
		j := getMid(nums[i])
		if j >= len(list) {
			list = append(list, nums[i])
		} else {
			list[j] = nums[i]
		}

		if len(list) > max {
			max = len(list)
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{-1}))                         //1
	fmt.Println(lengthOfLIS([]int{-2, -1}))                     //2
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})) //4
	fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})) //6
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))          // 3
	fmt.Println(lengthOfLIS([]int{2, 2}))                       // 1
}
