package main

func smallestSubsequence(s string) string {
	// 将元素依次入栈
	// 如果栈顶元素大于当前元素,且还会在出现,就移除

	// 记录每个字母最后一次出现的位置
	lastPosition := make(map[byte]int, len(s))

	// 记录这个字母有没有出现过
	visited := make(map[byte]bool, len(s))

	// 栈
	list := make([]byte, 0, len(s))

	// 记录每个字母最后一次出现的位置
	for i := 0; i < len(s); i++ {
		lastPosition[s[i]] = i
	}

	//
	for i := 0; i < len(s); i++ {
		j := len(list) - 1
		// 出现过
		if visited[s[i]] {
			continue
		}
		for ; j >= 0; j-- {
			if list[j] < s[i] || lastPosition[list[j]] < i {
				break
			}
			// 被移除了
			visited[list[j]] = false
		}
		list = list[:j+1]

		list = append(list, s[i])
		// 记录
		visited[s[i]] = true
	}
	return string(list)
}

func main() {

}
