package main

import "fmt"

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

func twoSum(n int) []float64 {
	// 动态规划
	// 状态定义: d[i][j] 投掷第i骰子时的和
	// 方程: d[i][j] = d[i-1][j-1] + d[i-1][j-2] + ... + d[i-1][j-6]
	// 边界 d[0][0] = 0, d[0][1] = 2, ... d[0][6] = 6
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n*6+6+1)
	}
	for i := 1; i <= 6; i++ {
		d[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= i*6+6; j++ {
			for k := 1; k <= 6 && j-k > 0; k++ {
				// fmt.Println(j, k, j-k.len(d[i-1]))
				d[i][j] += d[i-1][j-k]
			}
			fmt.Println(i, j, d[i][j])
		}
	}
	total := float64(0)
	for i := 1; i < len(d[n-1]); i++ {
		total += float64(d[n-1][i])
	}
	result := make([]float64, 0, len(d[n-1])-1)
	for i := 1; i < len(d[n-1]); i++ {
		rate := float64(d[n-1][i]) / total
		if rate > 0 {
			result = append(result, rate)
		}
	}
	return result
}

func main() {
	// fmt.Println(twoSum(1))
	fmt.Println(twoSum(2))
}