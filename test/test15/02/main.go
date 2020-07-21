package main

import "sort"

// 排序
func maxWidthRamp(A []int) int {
	// 将下标按照实际值排序
	indexArr := make([]int, len(A))
	for i, _ := range A {
		indexArr[i] = i
	}
	// 得是稳定排序算法...
	sort.Slice(indexArr, func(i, j int) bool {
		return A[i] < A[j]
	})

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxValue, minIndex := 0, len(A)
	for _, v := range indexArr {
		maxValue = max(maxValue, v-minIndex)
		minIndex = min(minIndex, v)
	}
	return maxValue
}

func combine(l1, l2 []int, less func(i, j int) bool) []int {
	l3 := make([]int, 0, len(l1)+len(l2))
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if less(i, j) {
			l3 = append(l3, l1[i])
			i++
		} else {
			l3 = append(l3, l2[j])
			j++
		}
	}
	if i < len(l1) {
		l3 = append(l3, l1[i:]...)
	}
	if j < len(l2) {
		l3 = append(l3, l2[j:]...)
	}
	return l3
}

//
func maxWidthRamp1(A []int) int {
	max := 0
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if A[j] <= A[i] {
				if i-j > max {
					max = i - j
				}
				break
			}
		}
	}

	return max
}

/*

[9,8,7,7,7,1,0,1,9,4,0,4,7,2]

*/
