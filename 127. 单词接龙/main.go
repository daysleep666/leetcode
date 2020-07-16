package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	// 找到相邻的单词
	isXL := func(a, b string) bool {
		c := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				c++
			}
			if c > 1 {
				return false
			}
		}
		return true
	}
	m := make(map[string]bool, len(wordList))
	for _, v := range wordList {
		m[v] = true
	}
	if !m[endWord] {
		return 0
	}

	find := func(str string) []string {
		arr := make([]string, 0)
		for v, _ := range m {
			if isXL(str, v) {
				delete(m, v)
				// fmt.Println(str, v)
				arr = append(arr, v)
			}
		}
		return arr
	}

	c := 1
	for len(queue) != 0 {
		c++
		tmp := make([]string, 0)
		for _, v := range queue {
			if isXL(endWord, v) {
				return c
			}
			a := find(v)
			tmp = append(tmp, a...)
		}
		queue = tmp
	}
	return 0
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
}

// "hit"
// "cog"
// ["hot","dot","dog","lot","log"]
