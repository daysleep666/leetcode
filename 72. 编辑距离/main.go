package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	min := func(list ...int) int {
		min := list[0]
		for _, v := range list {
			if v < min {
				min = v
			}
		}
		return min
	}
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	// DB
	// 状态定义: dp[i][j] 将0到i个字符转换为0到j个字符所花费的步数
	// 状态转移方程: dp[i][j] = min(dp[i-1][j-1], dp[i][j-1],dp[i-1][j]) + 1 // 替换,增加,删除
	// 如果word[i] == word[j] op[i][j] = op[i-1][j-1]
	op := make([][]int, len(word1))
	for i := 0; i < len(op); i++ {
		op[i] = make([]int, len(word2))
	}

	get := func(i, j int) int {
		// word1为空字符串, 那步数必然是word2
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		return op[i][j]
	}

	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				op[i][j] = get(i-1, j-1)
			} else {
				op[i][j] = min(get(i-1, j-1), get(i, j-1), get(i-1, j)) + 1
			}
		}
	}

	return op[len(word1)-1][len(word2)-1]
}

func minDistance1(word1 string, word2 string) int {
	log := map[[2]string]int{}
	return recursice(word1, word2, log)
}

func recursice(word1 string, word2 string, log map[[2]string]int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(list ...int) int {
		min := list[0]
		for _, v := range list {
			if v < min {
				min = v
			}
		}
		return min
	}
	if n, ok := log[[2]string{word1, word2}]; ok {
		return n
	}
	//
	if len(word1) == 0 || len(word2) == 0 {
		// 删除或添加
		return max(len(word1), len(word2))
	}

	// 如果最后一个字符相等 就舍去
	if word1[len(word1)-1] == word2[len(word2)-1] {
		return minDistance(word1[:len(word1)-1], word2[:len(word2)-1])
	}

	n := min(minDistance(word1, word2[:len(word2)-1]), // 相当于添加
		minDistance(word1[:len(word1)-1], word2),                // 相当于删除
		minDistance(word1[:len(word1)-1], word2[:len(word2)-1]), // 相当于替换
	) + 1
	log[[2]string{word1, word2}] = n
	return n
}

func main() {
	fmt.Println(minDistance("horse", "ros"))           // 3
	fmt.Println(minDistance("a", "b"))                 // 1
	fmt.Println(minDistance("ab", "bb"))               // 1
	fmt.Println(minDistance("ab", ""))                 // 2
	fmt.Println(minDistance("intention", "execution")) // 5
	fmt.Println(minDistance("ab", "a"))                // 1
	fmt.Println(minDistance("bc", "cd"))               // 2

}
