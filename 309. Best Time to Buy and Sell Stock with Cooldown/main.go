package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, 2)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][0][1])
		dp[i][0][1] = dp[i-1][1][0] + prices[i-1]
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][0][0]-prices[i-1])
	}
	return max(dp[len(dp)-1][0][0], dp[len(dp)-1][0][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{2, 1}))
}
