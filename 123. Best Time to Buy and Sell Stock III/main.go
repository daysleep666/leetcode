package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 5)
	}
	dp[0][1] = -prices[0]
	result := 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		dp[i][1] = getMax(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = getMax(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = getMax(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = getMax(dp[i-1][4], dp[i-1][3]+prices[i])
		result = getMax(getMax(result, dp[i][4]), dp[i][2])

	}
	fmt.Println(dp)
	return result
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
