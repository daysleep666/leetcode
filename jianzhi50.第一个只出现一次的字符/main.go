package main

func firstUniqChar1(s string) byte {
	// 记录每个字符出现次数
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 找到第一个出现一次的
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}

// 用数组的效率远高于map
func firstUniqChar(s string) byte {
	// 记录每个字符出现次数
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	// 找到第一个出现一次的
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {

}
