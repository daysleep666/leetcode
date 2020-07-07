package main

func findLength2(A []int, B []int) int {
	// 动态规划
	// d[i][j] A从i和B从j时相等的数量
	// d[i][j] = d[i-1][j-1] + 1
	d := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		d[i] = make([]int, len(B))
	}
	get := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return d[i][j]
	}
	result := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				d[i][j] = get(i-1, j-1) + 1
				result = max(result, d[i][j])
			}
		}
	}
	return result
}

func findLength1(A []int, B []int) int {
	return recursive(A, B, 0, 0)
}

func recursive(A, B []int, i, j int) int {
	if i == len(A) || j == len(B) {
		return 0
	}

	c := 0
	for m, n := i, j; m < len(A) && n < len(B); m, n = m+1, n+1 {
		if A[m] == B[n] {
			c++
		}
	}
	if c == len(A[i:]) || c == len(B[j:]) {
		return c
	}

	return max(recursive(A, B, i+1, j), recursive(A, B, i, j+1), c)
}

func max(nums ...int) int {
	m := 0
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}

func main() {
}
