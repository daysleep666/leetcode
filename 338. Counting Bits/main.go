package main

func countBits1(num int) []int {
	count := func(n int) (c int) {
		for n != 0 {
			n = n & (n - 1)
			c++
		}
		return
	}
	list := make([]int, num+1)
	for i := 0; i <= num; i++ {
		list[i] = count(i)
	}
	return list
}

// 这个方法很赞 去掉最低位的1,然后用这个数的1的数量+1 就是结果了
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}
