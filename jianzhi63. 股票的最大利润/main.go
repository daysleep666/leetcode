package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 动态规划
	// d[i][j] 第i天持有/不持有股票的资产
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	d := make([][2]int, len(prices))
	d[0][1] = -prices[0]
	m := 0
	for i := 1; i < len(prices); i++ {
		d[i][0] = d[i-1][1] + prices[i]
		d[i][1] = max(d[i-1][1], -prices[i])
		m = max(m, d[i][0])
	}
	return m
}

func main() {

}
