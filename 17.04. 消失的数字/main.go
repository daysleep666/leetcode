package main

// 两个相同的数异或的结果为0
func missingNumber(nums []int) int {
	bitmask := 0
	for i, v := range nums {
		bitmask ^= i
		bitmask ^= v
	}
	bitmask ^= len(nums)
	return bitmask
}

func main() {

}

// 00000000
// 00000010
// 00000100
// 00001000
// 00010000
// 00100000
// 01000000
// 10000000
// 10000001
// 10000010
// 10000100
// 10001000
