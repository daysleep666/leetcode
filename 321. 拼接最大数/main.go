package main

import "fmt"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var m []int
	// num1里选x个数, num2里选k-x个数. 归并
	for i := 0; i <= k && i <= len(nums1); i++ {
		a := getMax(nums1, len(nums1)-i)
		b := getMax(nums2, len(nums2)-(k-i))
		c := merge(a, b)
		m = max(m, c)
	}
	return m
}

// 移除k个元素,保持最大
func getMax(nums []int, k int) []int {
	if len(nums) <= k {
		return nil
	}
	if k <= 0 {
		return nums
	}
	// 如果栈顶元素小于当前元素就移除
	list := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		j := len(list) - 1
		for ; j >= 0; j-- {
			if list[j] >= nums[i] || k == 0 {
				break
			}
			k--
		}
		list = list[:j+1]
		list = append(list, nums[i])
	}

	// 如果k不等于0,就从尾部开始移除
	list = list[:len(list)-k]
	return list
}

func merge(l1, l2 []int) []int {
	compare := func(l1, l2 []int) bool {
		for i := 0; i < len(l1) && i < len(l2); i++ {
			if l1[i] < l2[i] {
				return false
			} else if l1[i] > l2[i] {
				return true
			}
		}
		return len(l1) > len(l2)
	}
	l := make([]int, 0, len(l1)+len(l2))
	for len(l1) > 0 && len(l2) > 0 {
		if compare(l1, l2) {
			l = append(l, l1[0])
			l1 = l1[1:]
		} else {
			l = append(l, l2[0])
			l2 = l2[1:]
		}
	}
	if len(l1) > 0 {
		l = append(l, l1...)
	}
	if len(l2) > 0 {
		l = append(l, l2...)
	}
	return l
}

func max(l1, l2 []int) []int {
	if len(l1) > len(l2) {
		return l1
	} else if len(l1) < len(l2) {
		return l2
	}
	for i := 0; i < len(l1); i++ {
		if l1[i] > l2[i] {
			return l1
		} else if l1[i] < l2[i] {
			return l2
		}
	}
	return l1
}

func main() {
	// fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3))
	// fmt.Println(maxNumber([]int{5, 5, 1}, []int{4, 0, 1}, 3))
	fmt.Println(maxNumber([]int{2, 5, 6, 4, 4, 0}, []int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15))
	fmt.Println(maxNumber([]int{0}, []int{0, 6, 5, 7, 6, 2}, 15))

}

// ]
// []
// 15

// [7,3,8,2,5,6,4,4,0,6,5,7,6,2,0]
