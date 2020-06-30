package main

func findLUSlength(strs []string) int {
	list := make([]string, 0, len(strs))
	list = append(list, strs[0])

	for i := 1; i < len(strs); i++ {
		j := len(list) - 1
		for ; j >= 0; j-- {
			if len(strs[i]) == len(list[j]) {
				if equal(strs[i], list[j]) {
					list = append(list[:j], list[j+1:]...)
				} else {
					list = append(list[:j+1], strs[i], list[j+1])
				}
				break
			} else if len(strs[i]) < len(list[j]) {
				if isSub(list[j], strs[i]) {
					break
				}
			} else {
				if isSub(list[j], strs[i]) {
					list[j] = strs[i]
				} else {
					list = append(list[:j+1], strs[i], list[j+1:])
				}
				break
			}
			if j < 0 {
				list = append(list[:0], strs[i], list[1:])
			}
		}
	}
	if len(list) == 0 {

	}
}

func equal(a, b string) bool {
	return a == b
}

func isSub(a, b string) bool {
	if len(a) < len(b) {
		a, b = b, a
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			j++
		}
		i++
	}

	return j == len(b)
}

func main() {

}

// 给定字符串列表，你需要从它们中找出最长的特殊序列。最长特殊序列定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。

// 子序列可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。

// 输入将是一个字符串列表，输出是最长特殊序列的长度。如果最长特殊序列不存在，返回 -1 。

// 示例：

// 输入: "aba", "cdc", "eae"
// 输出: 3

// 提示：

// 所有给定的字符串长度不会超过 10 。
// 给定字符串列表的长度将在 [2, 50 ] 之间。
