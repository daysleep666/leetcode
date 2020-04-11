package main

import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	dp := make([]int, amount+1) // dp[i]到第i阶的最小步数
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 暴力求解... 时间复杂度O(2^n) TODO: 在查下
func coinChange1(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	sort.Ints(coins)

	return recursive(coins, amount, map[int]int{})
}

func recursive(coins []int, amount int, m map[int]int) int {
	min := math.MaxInt32
	for i := len(coins) - 1; i >= 0; i-- {
		if amount-coins[i] == 0 {
			return 1
		} else if amount-coins[i] < 0 {
			continue
		} else {
			r, ok := 0, false
			if r, ok = m[amount-coins[i]]; !ok {
				r = recursive(coins, amount-coins[i], m)
				m[amount-coins[i]] = r
			}
			if r == -1 {
				continue
			}
			if r < min {
				min = r
			}
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min + 1
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))             // 3
	fmt.Println(coinChange([]int{2}, 11))                   // -1
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))         // 4
	fmt.Println(coinChange([]int{2, 7}, 10))                // 5
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249)) // 20
}
