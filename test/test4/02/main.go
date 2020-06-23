package main

// 位运算
func singleNumber(nums []int) int {
	var once, twice int
	for _, v := range nums {
		// 出现一次 在once里
		// 出现两次 在twice里
		// 出现三次 没了
		once = (^twice) & (once ^ v)
		twice = (^once) & (twice ^ v)
		// once= (~twice)&(once^v)
	}
	return once
}

func main() {

}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

// 说明：

// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// 示例 1:

// 输入: [2,2,3,2]
// 输出: 3
// 示例 2:

// 输入: [0,1,0,1,0,1,99]
// 输出: 99
