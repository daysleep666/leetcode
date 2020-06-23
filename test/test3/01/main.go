package main

import "fmt"

// 时间复杂度: log(n*m*m)
// 空间复杂度: log(m+n)
func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	bStr1 := []byte(str1)
	old := []byte(str2)
	list := make([]byte, len(old))
	// 保证str2<=str1
	for i := 0; i < len(str2); i++ {
		if len(old)%len(old[i:]) != 0 {
			continue
		}
		copy(list[len(old)-i:], old[:i])
		copy(list[:len(old)-i], old[i:])
		if compare(old, list) {
			if isRepeat(bStr1, old[i:]) {
				return string(old[i:])
			}
		}
	}
	return ""
}

// repstr是否是str的循环子字符串
func isRepeat(str, repStr []byte) bool {
	if len(str)%len(repStr) != 0 {
		return false
	}
	for i := 0; i < len(str); i++ {
		if str[i] != repStr[i%len(repStr)] {
			return false
		}
	}
	return true
}

func compare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(gcdOfStrings("ABC", "ABC"))     // ABC
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // ABC
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // AB
	fmt.Println(gcdOfStrings("LEET", "ABAB"))   // ""
	fmt.Println(gcdOfStrings("AAAAAA", "AAAA")) // "AA"
}

// 对于字符串 S 和 T，只有在 S = T + ... + T（T 与自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

// 返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。

// 示例 1：

// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
// 示例 2：

// 输入：str1 = "ABABAB", str2 = "ABAB"
// 输出："AB"
// 示例 3：

// 输入：str1 = "LEET", str2 = "CODE"
// 输出：""

// 提示：

// 1 <= str1.length <= 1000
// 1 <= str2.length <= 1000
// str1[i] 和 str2[i] 为大写英文字母
