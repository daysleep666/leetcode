package main

import "fmt"

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(3))
	fmt.Println(integerBreak(6))
}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	// 动态规划
	// dp[i] 当i时的最大乘积
	// dp[i] = max(dp[i],max(d[i-j],d[i-j]*(i-j)))
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], dp[i-j]*j)
			// fmt.Println(i, j, dp[i])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
