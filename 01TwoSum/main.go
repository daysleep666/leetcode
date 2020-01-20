package main

import "fmt"



// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum(nums []int, target int) []int {
	find := make(map[int]int)
	for i, v := range nums {
		if j, ok := find[target-v]; ok {
			return []int{j, i}
		}
		find[v] = i
	}

	return nil
}

func main() {
	{
		result := twoSum([]int{2, 7, 11, 15}, 9)
		fmt.Println(result)
	}
	{
		result := twoSum([]int{3, 2, 4}, 6)
		fmt.Println(result)
	}
}

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].
