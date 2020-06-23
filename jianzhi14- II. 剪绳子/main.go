package main

import "fmt"

func cuttingRope(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	// 动态规划
	// 状态定义: d[i][j] 长度为i时分成j段的最大乘积
	// 转移方程: d[i][j] = max(d[i-x][j-1]*x) (x:1-(i-1))
	d := make([][]int, n+1)
	for i := 0; i < len(d); i++ {
		d[i] = make([]int, n+1)
	}
	result := 0
	for i := 3; i <= n; i++ {
		for j := 2; j < n; j++ {
			max := 0
			for x := 1; x < i-1; x++ {
				d[i-x][1] = i - x
				if d[i-x][j-1]*x > max {
					// TODO: 剪枝
					max = d[i-x][j-1] * x
					fmt.Println((d[i-x][j-1] * x) % 1000000007)
				}
			}
			if max > result {
				result = max
			}
			d[i][j] = max
		}
	}
	return result % 1000000007
}

func main() {
	// fmt.Println(cuttingRope(3)) //
	// fmt.Println(cuttingRope(4)) //
	// fmt.Println(cuttingRope(5))   //
	// fmt.Println(cuttingRope(6))   //
	// fmt.Println(cuttingRope(7))   //
	// fmt.Println(cuttingRope(8))   //
	// fmt.Println(cuttingRope(9))   //
	// fmt.Println(cuttingRope(10))  //
	fmt.Println(cuttingRope(120)) //
}
