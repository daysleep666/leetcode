package main

func singleNumbers(nums []int) []int {
	// 依次异或获得的结果就是只出现单次

	numbers(nums[:len(nums)/2])
	numbers(nums[len(nums)/2:])
}

func numbers(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum ^= nums[i]
	}
	return sum
}
