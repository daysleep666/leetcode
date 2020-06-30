package main

func nextPermutation(nums []int) {
	// 从后往前 如果是降序就调换位置
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			// 找到右侧刚好比它大一位的交换
			min := i
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < nums[min] && nums[j] > nums[i-1] {
					min = j
				}
			}
			nums[i-1], nums[min] = nums[min], nums[i-1]

			// 将右侧变为升序
			l := sort(nums[i:])
			copy(nums[i:], l)
			break
		}
	}
	// 说明一直是升序 也就是最大了
	if i == 0 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	a := sort(arr[:len(arr)/2])
	b := sort(arr[len(arr)/2:])

	list := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			list = append(list, a[i])
			i++
		} else {
			list = append(list, b[j])
			j++
		}
	}
	if i < len(a) {
		list = append(list, a[i:]...)
	}
	if j < len(b) {
		list = append(list, b[j:]...)
	}
	return list
}

func main() {

}
