package main

// 双指针
func reverseVowels(s string) string {
	isVowel := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
			b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
	}
	start, end := 0, len(s)-1
	list := []byte(s)
	for start < end {
		if !isVowel(list[start]) {
			start++
		}
		if !isVowel(list[end]) {
			end--
		}
		if isVowel(list[start]) && isVowel(list[end]) {
			list[start], list[end] = list[end], list[start]
			start++
			end--
		}
	}
	return string(list)
}

// //
// 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

// 示例 1:

// 输入: "hello"
// 输出: "holle"
// 示例 2:

// 输入: "leetcode"
// 输出: "leotcede"
// 说明:
// 元音字母不包含字母"y"。
