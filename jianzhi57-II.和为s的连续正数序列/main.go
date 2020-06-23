package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)
	sum := 0
	list := make([]int, 0, target)
	for i := 1; i <= target/2+1; i++ {
		list = append(list, i)
		sum += i
		for sum > target && len(list) > 0 {
			sum -= list[0]
			list = list[1:]
		}
		if len(list) == 0 {
			break
		}
		if sum == target {
			tmp := make([]int, len(list))
			copy(tmp, list)
			result = append(result, tmp)
		}
	}
	return result
}

func main() {
	fmt.Println(findContinuousSequence(87760))
}
