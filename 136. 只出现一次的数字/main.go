package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func main() {
	fmt.Println(0 ^ 8)
}

// 异或

// 相同是0 不相同是1
