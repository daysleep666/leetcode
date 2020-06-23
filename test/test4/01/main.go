package main

import (
	"fmt"
)

// 比较两个数组合后的大小: 组合后直接比较!!!

// 排序
func largestNumber(nums []int) string {
	// 将[]int变为[][]byte
	arr := make([][]byte, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[i] = int2Bytes(nums[i])
	}
	quickSort(arr, 0, len(nums)-1)
	result := make([]byte, 0, len(arr)*2)
	for _, v := range arr {
		result = append(result, v...)
	}
	i := 0
	for ; i < len(result); i++ {
		if result[i] != '0' {
			break
		}
	}
	result = result[i:]
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}

// 大到小排序
func quickSort(num [][]byte, start, end int) {
	if start >= end {
		return
	}
	mid := partition(num, start, end)
	// fmt.Println(num, mid)
	quickSort(num, start, mid-1)
	quickSort(num, mid+1, end)
}

func partition(num [][]byte, start, end int) int {
	i, j := start, start
	for ; i < end; i++ {
		if greater(num[i], num[end]) {
			num[i], num[j] = num[j], num[i]
			j++
		}
	}
	num[j], num[end] = num[end], num[j]
	return j
}

func greater(a, b []byte) bool {

	m := make([]byte, 0, len(a)+len(b))
	n := make([]byte, 0, len(a)+len(b))
	m = append(m, a...)
	m = append(m, b...)
	n = append(n, b...)
	n = append(n, a...)
	for i := 0; i < len(m); i++ {
		if m[i] > n[i] {
			return true
		} else if m[i] < n[i] {
			return false
		}
	}
	return true
}

func int2Bytes(n int) []byte {
	if n == 0 {
		return []byte{'0'}
	}
	arr := make([]byte, 0, 100)
	for n != 0 {
		a := n % 10
		arr = append(arr, byte(a+'0'))
		n = n / 10
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	// fmt.Println(largestNumber([]int{1, 2, 3}))
	// fmt.Println(largestNumber([]int{10, 2}))
	// fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	// fmt.Println(largestNumber([]int{0, 1}))
	fmt.Println(largestNumber([]int{121, 12}))
}

// 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

// 示例 1:

// 输入: [10,2]
// 输出: 210
// 示例 2:

// 输入: [3,30,34,5,9]
// 输出: 9534330
// 说明: 输出结果可能非常大，所以你需要返回一个字符串而不是整数。
