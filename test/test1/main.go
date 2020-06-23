package main

import "fmt"

func sequentialDigits(low int, high int) []int {
	result := make([]int, 0)
	i := 0
	c := 1

	for i <= high {
		i = c
		c = (c + 1) % 10
		arr := make([]int, 0)
		for {
			arr = append(arr, i%10+1)
			i = i*10 + (i%10 + 1)
			if i >= low {
				break
			}
		}
		if i < high {
			result = append(result, i)
		}
		low = i + 1
	}
	return result
}

func main() {
	fmt.Println(sequentialDigits(100, 200))
	fmt.Println(sequentialDigits(1000, 13000))
}

// 我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。

// 请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）。

// 示例 1：

// 输出：low = 100, high = 300
// 输出：[123,234]
// 示例 2：

// 输出：low = 1000, high = 13000
// 输出：[1234,2345,3456,4567,5678,6789,12345]

// 提示：

// 10 <= low <= high <= 10^9
