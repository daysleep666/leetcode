package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	/*第n天*/ /*0没有股票1持有1股*/ /*0正常1冷却*/
	dp := make([][][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([][]int, 2)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	// dp[0][1][0] = -prices[i]
	for i := 1; i < len(prices); i++ {
		dp[i][1][0] = getMax(dp[i-1][1][0], dp[i-1][0][0]-prices[i])
		dp[i][0][1] = dp[i-1][1][0] + prices[i]
		dp[i][0][0] = getMax(dp[i-1][0][1], dp[i-1][0][0])
	}
	return getMax(dp[len(dp)-1][0][0], dp[len(dp)-1][0][1])
}

func main() {
	fmt.Println(maxProfit([]int{2, 1}))
}
