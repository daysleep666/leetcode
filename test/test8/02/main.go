package main

import "fmt"

// 哈希表
// 堆
func topKFrequent(words []string, k int) []string {
	// 计算每个单词出现的次数
	m := make(map[string]int, len(words))
	arr := make([]string, 0, len(words))
	for _, v := range words {
		if _, ok := m[v]; !ok {
			arr = append(arr, v)
		}
		m[v]++
	}

	// heap 大顶堆
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, m, i, len(arr))
	}

	result := make([]string, k)
	// 堆排序
	for i, end := 0, len(arr)-1; i < k; i, end = i+1, end-1 {
		result[i] = arr[0]
		arr[0], arr[end] = arr[end], arr[0]
		heapify(arr, m, 0, end)
	}
	return result
}

func heapify(arr []string, m map[string]int, start, end int) {
	cur := start

	if start*2+1 < end && (m[arr[start*2+1]] > m[arr[cur]] || (m[arr[start*2+1]] == m[arr[cur]] && arr[start*2+1] < arr[cur])) {
		cur = start*2 + 1
	}
	if start*2+2 < end && (m[arr[start*2+2]] > m[arr[cur]] || (m[arr[start*2+2]] == m[arr[cur]] && arr[start*2+2] < arr[cur])) {
		cur = start*2 + 2
	}
	if cur == start {
		return
	}
	arr[cur], arr[start] = arr[start], arr[cur]
	heapify(arr, m, cur, end)
}

func main() {
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3))

}
