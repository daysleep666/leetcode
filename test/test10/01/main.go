package main

import "sort"

// 贪心
// 用最小的饼干满足胃口最小的孩子
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for i := 0; i < len(s) && count < len(g); i++ {
		if s[i] >= g[count] {
			count++
		}
	}
	return count
}

func main() {

}
