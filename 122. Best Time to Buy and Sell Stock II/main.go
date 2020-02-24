package main

import "fmt"

func maxProfit1(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			dp[i] = dp[i-1] + prices[i] - prices[i-1]
			continue
		}
		dp[i] = dp[i-1]
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
