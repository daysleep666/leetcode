package main

// 蠢方法
func hammingWeight1(num uint32) int {
	var result int
	for num != 0 {
		if num&1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

// x&(x-1) // 可以消除掉末尾的1
// 例子 1110&1101=1100   1111&1110=1110
func hammingWeight(num uint32) int {
	var result int
	for num != 0 {
		num = num & (num - 1)
		result++
	}
	return result
}

func main() {

}
