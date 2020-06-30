package main

import "sort"

// 贪心
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	end := points[0][1]
	arrows := 1
	for i := 1; i < len(points); i++ {
		p := points[i]
		if p[0] <= end {
			continue
		}
		arrows++
		end = p[1]
	}
	return arrows
}

// 排序

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
