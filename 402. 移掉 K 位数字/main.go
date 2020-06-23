package main

// 移除一个数组里的k位,保持相对位置不变
// 栈 + 贪心

func removeKdigits(num string, k int) string {
	if len(num) == 0 {
		return "0"
	}

	list := make([]byte, 0, len(num))
	// 如果栈顶元素大于当前元素,就移除它,
	// 如果k==len(num[i:]) 就直接移除
	for i := 0; i < len(num); i++ {
		j := len(list) - 1
		for ; j >= 0 && k > 0; j-- {
			if list[j] <= num[i] {
				break
			}
			k--
		}
		list = list[:j+1]
		list = append(list, num[i])
	}
	if k != 0 {
		list = list[:len(list)-k]
	}
	i := 0
	for ; i < len(list); i++ {
		if list[i] != '0' {
			break
		}
	}
	list = list[i:]
	if len(list) == 0 {
		return "0"
	}
	return string(list)
}

func main() {

}
