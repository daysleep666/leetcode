package main

import "fmt"

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int, len(words))
	arr := make([]string, 0, len(words))
	for _, w := range words {
		_, ok := m[w]
		if !ok {
			arr = append(arr, w)
		}
		m[w]++

	}
	// 堆化--大顶堆
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, m, i, len(arr))
	}
	l := len(arr)
	for i := 0; i < k && l > 1; i++ {
		arr[0], arr[l-1] = arr[l-1], arr[0]
		heapify(arr, m, 0, l-1)
		l--
	}
	// reverse
	for i, j := 0, len(arr)-1; i < j && i < k; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr[:k]
}

func heapify(arr []string, m map[string]int, i int, l int) {
	big := func(i, j int) bool {
		if m[arr[i]] == m[arr[j]] {
			return arr[i] < arr[j]
		}
		return m[arr[i]] > m[arr[j]]
	}
	n := i
	if i*2+1 < l && big(i*2+1, i) {
		n = i*2 + 1
	}
	if i*2+2 < l && big(i*2+2, n) {
		n = i*2 + 2
	}
	if n == i {
		return
	}
	arr[i], arr[n] = arr[n], arr[i]
	heapify(arr, m, n, l)
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
