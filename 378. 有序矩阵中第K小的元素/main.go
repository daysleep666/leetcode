package main

func kthSmallest(matrix [][]int, k int) int {
	// 二分 这思路...

}

// 计算矩阵中小于等于mid的个数
func check(matrix [][]int, mid, k int) bool {
	num := 0
	for i, j := 0, 0; i < len(matrix) && 0 <= j && k < len(matrix[i]); {
		if matrix[i][j] <= mid {
			j++
		} else {
			num += j - 1
			i++
			j--
		}
	}
	return num >= k
}

func kthSmallest1(matrix [][]int, k int) int {
	l := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		l = combine(l, matrix[i])
	}
	return l[k-1]
}

func combine(l1, l2 []int) []int {
	l3 := make([]int, 0, len(l1)+len(l2))
	i, j := 0, 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
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

func main() {

}
