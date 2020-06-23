package main

func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{nums[i], nums[j]}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
