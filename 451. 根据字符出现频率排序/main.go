package main

import (
	"fmt"
)

func frequencySort(s string) string {
	var m = make(map[rune]int, len(s))
	arr := make([]rune, 0, len(s))
	for _, v := range s {
		if _, ok := m[v]; !ok {
			arr = append(arr, v)
		}
		m[v]++
	}
	quickSort(arr, 0, len(arr)-1, func(i, j int) bool {
		return m[arr[i]] > m[arr[j]]
	})

	result := make([]rune, 0, len(s))
	for _, r := range arr {
		num := m[r]
		for i := 0; i < num; i++ {
			result = append(result, r)
		}
	}

	return string(result)
}

func quickSort(arr []rune, start, end int, f func(i, j int) bool) {
	if len(arr) <= 1 {
		return
	}
	if start >= end {
		return
	}
	mid := partition(arr, start, end, f)
	quickSort(arr, start, mid-1, f)
	quickSort(arr, mid+1, end, f)
	return
}

func partition(arr []rune, start, end int, f func(i, j int) bool) int {
	i, j := start, start

	for i < end {
		if f(i, end) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
			continue
		}
		i++
	}
	arr[j], arr[end] = arr[end], arr[j]
	return j
}

func p(arr []rune) {
	for _, v := range arr {
		fmt.Printf("%s ", string(v))
	}
	fmt.Println("-----")
}

func main() {
	str := "acbbecefefeffcf"
	// str := "2a552f544afaffffa"
	// for i := 0; i < 10000; i++ {
	// 	str += string(rune(rand.Int63n(25) + 'a'))
	// }
	// fmt.Println(str)
	fmt.Println(frequencySort(str))
}
