package main

// 双指针
func reverseWords(s string) string {
	// 1. 数组存放words
	arr := make([]string, 0)
	// 2. start,end 标记单词上下界
	start, end := 0, 0
	// 3. 确认边界
	for ; end < len(s); end++ {
		// 如果是空格. 则说明[start, end]是一个单词
		if s[end] == ' ' {
			if start != end {
				arr = append(arr, s[start:end])
			}
			start = end + 1
		}
	}
	if start != end {
		arr = append(arr, s[start:end])
	}
	s = ""
	for i := len(arr) - 1; i >= 0; i-- {
		s += arr[i]
		if i != 0 {
			s += " "
		}
	}
	return s
}
