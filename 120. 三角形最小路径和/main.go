package main

import "math"

// 动态规划
func minimumTotal(triangle [][]int) int {
	// dp[i] 第i个节点的最小和
	// dp[i] = 上一层 min(dp[i], dp[i-1])
	get := func(i, j int) int {
		if i < 0 || j < 0 || j >= len(triangle[i]) {
			return int(math.MaxInt64)
		}
		return triangle[i][j]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = min(get(i-1, j), get(i-1, j-1)) + triangle[i][j]
		}
	}
	result := int(math.MaxInt64)
	for i, j := len(triangle)-1, 0; j < len(triangle[i]); j++ {
		result = min(result, triangle[i][j])
	}
	return result
}
