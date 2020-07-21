package main

// 动态规划
func constrainedSubsetSum(nums []int, k int) int {
	// d[i] 到第i时的最大值
	// 动态转移方程: d[i] = nums[i] + max(0, d[i-1],d[i-2]...d[i-k])
	max := func(arr ...int) int {
		m := arr[0]
		for _, v := range arr {
			if v > m {
				m = v
			}
		}
		return m
	}
	maxValue := nums[0]
	d := make([]int, len(nums))
	d[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax := 0
		for j := i - 1; j < i && j >= 0; j-- {
			tmpMax = max(tmpMax, d[j])
		}
		d[i] = max(nums[i]+tmpMax, tmpMax)
		maxValue = max(maxValue, d[i])
	}
	return maxValue
}
