package main

import "fmt"

// 时间复杂度 x*n^2 x是平均字符串长度
// 空间复杂度 n
func findLUSlength(strs []string) int {
	visited := make(map[string]int, len(strs))
	for i := 0; i < len(strs); i++ {
		visited[strs[i]] = len(strs[i])
	}
	for i := 0; i < len(strs); i++ {
		if visited[strs[i]] == -1 {
			continue
		}
		for j := i + 1; j < len(strs); j++ {
			if visited[strs[j]] == -1 {
				continue
			}
			switch {
			case len(strs[i]) == len(strs[j]):
				if strs[i] == strs[j] {
					visited[strs[j]] = -1
				}

			case len(strs[i]) > len(strs[j]):
				if isSub(strs[i], strs[j]) {
					visited[strs[j]] = -1
				}

			case len(strs[i]) < len(strs[j]):
				if isSub(strs[j], strs[i]) {
					visited[strs[i]] = -1
					break
				}
			}

		}
	}
	max := -1
	for _, v := range visited {
		if v > max {
			max = v
		}
	}
	return max
}

func isSub(str string, subStr string) bool {
	i, j := 0, 0
	for ; i < len(str) && j < len(subStr); i++ {
		if str[i] == subStr[j] {
			j++
		}
	}
	return j == len(subStr)
}

func main() {
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "b"}))
}
