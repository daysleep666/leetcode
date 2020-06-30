package main

import "fmt"

// 时间复杂度 o(n)
func minCostClimbingStairs1(cost []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	cost = append(cost, 0)
	// 状态定义: 到第i阶台阶时的最低花费
	// 状态转移方程: dp[i] = min(dp[i-1],dp[i-2])
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}

	return dp[len(cost)-1]
}

func minCostClimbingStairs(cost []int) int {
	return min(recursive(cost, 0, 0), recursive(cost, 1, 0))
}

func recursive(cost []int, cur int, spend int) int {
	// 退出条件
	if cur >= len(cost) {
		return spend
	}

	// 走一步 or 走两步
	return min(recursive(cost, cur+1, spend+cost[cur]), recursive(cost, cur+2, spend+cost[cur]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 2, 3}))
}
