package main

import (
	"fmt"
)

func grayCode(n int) []int {
	list := []int{0}
	a := 1
	for i := 0; i < n; i++ {
		l := len(list)
		for j := l - 1; j >= 0; j-- {
			list = append(list, list[j]|a)
		}
		a = a << 1
	}
	return list
}

func main() {
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(3))
}

// 示例 1:

// 输入: 2
// 输出: [0,1,3,2]
// 解释:
// 00 - 0
// 01 - 1
// 11 - 3
// 10 - 2

// 对于给定的 n，其格雷编码序列并不唯一。
// 例如，[0,2,3,1] 也是一个有效的格雷编码序列。

// 00 - 0
// 10 - 2
// 11 - 3
// 01 - 1
// 示例 2:

// 输入: 0
// 输出: [0]
// 解释: 我们定义格雷编码序列必须以 0 开头。
//      给定编码总位数为 n 的格雷编码序列，其长度为 2n。当 n = 0 时，长度为 20 = 1。
//      因此，当 n = 0 时，其格雷编码序列为 [0]。

// 0000
// 0001

// 0011
// 0010

// 0110
// 0111
// 0101
// 0100

// 1100
// 1101
// 1111
// 1110
// 1010
// 1011
// 1001
// 1000
