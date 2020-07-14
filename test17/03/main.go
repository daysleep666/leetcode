package main

import "fmt"

// 动态规划
// 递归改动态规划。有几个可变参数决定你维护的dp数组是几维数组
func numWays(steps int, arrLen int) int {
	// 状态定义 dp[i][j] 第i步距离原点j时的方案数
	// 状态转移方程 dp[i][j] = dp[i-1][j] + dp[i-1][j+1] + dp[i-1][j-1]
	dp := make([][]int, steps+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, min(steps+1, arrLen))
	}
	dp[1][0] = 1
	dp[1][1] = 1
	get := func(i, j int) int {
		if j < 0 || j > i || j >= arrLen {
			return 0
		}
		// fmt.Println("-->", i, j, len(dp[i]))
		return dp[i][j]
	}
	for i := 2; i <= steps; i++ {
		for j := 0; j < arrLen && j <= i; j++ {
			dp[i][j] = (get(i-1, j) + get(i-1, j-1) + get(i-1, j+1)) % (1000000007)
		}
	}
	// fmt.Println("-->", dp[steps][0])
	return dp[steps][0] % (1000000007)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numWays(4, 2))
	fmt.Println(numWays(6, 13))
	fmt.Println(numWays(27, 7))
	fmt.Println(numWays(47, 38))
	fmt.Println(numWays(438, 315977))
}
