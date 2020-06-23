package main

import "fmt"

// 哈希表 借助哈希表去重
// 因为仅包含小写字母 所以使用[26]int
func removeDuplicateLetters(s string) string {
	var m [26]int
	for i := 0; i < 26; i++ {
		m[i] = -1
	}
	c := 0
	result := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		// 如果该字母没有出现过,就加入到结果集中,并标记
		b := s[i] - 'a'
		if m[b] != -1 {
			isReplace := false
			// 判断两次之间有没有小于当前值的,如果有就改为后者
			for j := m[b] + 1; j < len(result); j++ {
				if result[j] < s[i] {
					// 移除之前的那个,并添加新的
					isReplace = true
					break
				}
			}
			if isReplace {
				// 将抛弃的值替换为#
				result[m[b]] = '#'
			} else {
				continue
			}
		} else {
			c++
		}
		// 记录最近一次在result中出现的位置
		m[b] = len(result)
		result = append(result, s[i])
	}
	r := make([]byte, 0, len(result))
	for _, v := range result {
		if v != '#' {
			r = append(r, v)
		}
	}
	return string(r)
}

func main() {
	// fmt.Println(removeDuplicateLetters("acbc"))
	fmt.Println(removeDuplicateLetters("bcabc"))
}

// 给你一个仅包含小写字母的字符串，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

// 示例 1:

// 输入: "bcabc"
// 输出: "abc"
// 示例 2:

// 输入: "cbacdcbc"
// 输出: "acdb"
