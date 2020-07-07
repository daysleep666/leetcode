package main

// 1. 计算某个前缀的节点个数
// 2. 如果在当前前缀下,找到子节点
// 3. 如果不在当前前缀下,找到下一个前缀
func findKthNumber(n int, k int) int {
	prefix := 1
	// 已经跳过了多少个节点
	p := 1

	for p < k {
		// 计算以prefix开头的节点数量
		count := calPrefixCount(prefix, n)
		// 不在当前前缀下. 需要找到下一个节点
		if p+count <= k {
			prefix++
			p += count
		} else { // 在当前前缀下
			prefix *= 10
			p++
		}
	}
	return prefix
}

// 计算某个前缀的节点个数
// 下一个前缀的起始-当前前缀的起始
func calPrefixCount(prefix, n int) int {
	count := 0
	next := prefix + 1
	for prefix < n {
		// 因为n也算是节点之一 所以要额外加一
		count += min(next, n+1) - prefix
		prefix *= 10
		next *= 10
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
