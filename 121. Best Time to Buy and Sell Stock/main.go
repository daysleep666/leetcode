package main

import "fmt"

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	prices[0] = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			prices[i] = 0
		} else {
			prices[i] = prices[i] - min
		}
	}
	var max int
	for _, v := range prices {
		if v > max {
			max = v
		}
	}
	return max
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, result := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			if prices[i]-min > result {
				result = prices[i] - min
			}
		}
	}
	return result
}

func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMin := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = prices[0], 0
	for i := 1; i <= len(prices); i++ {
		dp[i][0] = getMin(dp[i-1][0], prices[i])
		dp[i][1] = getMax(dp[i-1][1], prices[i]-dp[i-1][0])
	}
	return dp[len(dp)-1][1]
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}
	k := 0
	for i := 1; i < len(prices); i++ {
		if (dp[i-1][1]+prices[i]) > dp[i-1][0] && k == 0 {
			dp[i][0] = dp[i-1][1] + prices[i]
		} else {
			dp[i][0] = dp[i-1][0]
		}
		if (dp[i-1][0]-prices[i] > dp[i-1][1]) && k == 0 {
			dp[i][1] = dp[i-1][0] - prices[i]
			k++
		} else {
			dp[i][1] = dp[i-1][1]
		}
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func main() {
}
