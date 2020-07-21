package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 4}))
	fmt.Println(rob([]int{1, 2, 3, 4}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(myRob(nums[1:]), myRob(nums[:len(nums)-1]))
}

func myRob(nums []int) int {
	// 动态规划
	// 状态定义: d[i][j] 第i间屋子 0不偷/1偷
	// 状态转移方程: d[i][j]
	//	d[i][0] = max(d[i-1][0], d[i-1][1])
	//	d[i][1] = d[i-1][0] + nums[i]
	// max(d[n][0],d[n][1])
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	d := make([][2]int, len(nums))
	d[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		d[i][0] = max(d[i-1][0], d[i-1][1])
		d[i][1] = d[i-1][0] + nums[i]
	}
	return max(d[len(d)-1][0], d[len(d)-1][1])
}

func max(nums ...int) int {
	m := 0
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
func rob1(nums []int) int {
	// 动态规划
	// 状态定义: d[i][j] 第i间屋子 0不偷/1偷
	// 状态转移方程: d[i][j]
	//	d[i][0] = max(d[i-1][0], d[i-1][1])
	//	d[i][1] = d[i-1][0] + nums[i]
	// max(d[n][0],d[n][1])

	d := make([][2][2]int, len(nums))
	d[0][0][0] = 0
	d[0][1][1] = nums[0]
	// d[1][0][0] = 0
	// d[1][1][0] = nums[1]
	// d[1][0][1] = d[0][1][1]
	for i := 1; i < len(nums); i++ {
		d[i][0][0] = max(d[i-1][0][0], d[i-1][1][0])
		d[i][1][0] = d[i-1][0][0] + nums[i]
		d[i][0][1] = max(d[i-1][0][1], d[i-1][1][1])
		d[i][1][1] = d[i-1][0][1] + nums[i]
	}
	return max(d[len(nums)-1][0][0], d[len(nums)-1][1][0], d[len(nums)-1][0][1])
}
