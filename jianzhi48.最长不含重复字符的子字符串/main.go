package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 动态规划
	// j是右边界 i是左边界
	// 状态定义 d[i]当前最长的不重复子字符串长度
	// 方程定义: d[i] = d[i-j] + (i-j)

	// 字母和下标的映射
	m := make(map[byte]int, len(s))
	// 最大值
	d := make([]int, len(s))
	i := 0
	// 左边界
	j := 0
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for ; i < len(s); i++ {
		// 如果访问到遍历过的节点 就更新左边界
		if v, ok := m[s[i]]; ok {
			j = getMax(j, v)
		}

		m[s[i]] = i

		d[i] = d[i-j] + (i - j)
	}

	max := 0
	for _, v := range d {
		if v > max {
			max = v
		}
	}
	return max
}

func lengthOfLongestSubstring1(s string) int {
	// 字母和下标的映射
	m := make(map[byte]int, len(s))
	// 最大值
	max := 0
	i := 0
	// 左边界
	j := -1
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for ; i < len(s); i++ {
		// 如果访问到遍历过的节点 就更新左边界
		if v, ok := m[s[i]]; ok {
			j = getMax(j, v)
		}

		m[s[i]] = i
		// 左右边界只差就是长度
		max = getMax(max, i-j)
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abca"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("ab"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring(""))
}
