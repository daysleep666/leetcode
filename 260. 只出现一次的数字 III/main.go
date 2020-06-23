package main

// 两个数字至少有一位不一样
// 求出不一样的那一位, 然后根据那一位将数组分成两组
func singleNumber(nums []int) []int {
	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}
	// 获取最右边的那个不一样的1
	diff := (^bitmask + 1) & bitmask
	// 按这一位是否是1来分组
	var a int
	for _, v := range nums {
		if diff&v == 0 {
			a ^= v
		}
	}
	// a^b = diff ---> a^diff=b
	return []int{a, diff ^ a}
}

func main() {

}
