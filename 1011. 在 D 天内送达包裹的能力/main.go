package main

import (
	"fmt"
)

func shipWithinDays1(weights []int, D int) int {
	max := 0
	for _, v := range weights {
		if v > max {
			max = v
		}
	}
	for {
		sum := 0
		d := 1
		for i, v := range weights {
			if len(weights[i:]) == D-d {
				return max
			}
			if sum+v <= max {
				sum += v
				continue
			}
			d++
			sum = v
		}
		if d == D {
			return max
		}
		max++
	}
}

func shipWithinDays(weights []int, D int) int {
	min, max := 0, 0
	for _, v := range weights {
		if v > min {
			min = v
		}
		max += v
	}

	can := func(w int) int {
		sum := 0
		d := 1
		for i, v := range weights {
			if len(weights[i:]) == D-d {
				return 0
			}
			if sum+v <= w {
				sum += v
				continue
			}
			d++
			sum = v
		}
		if d == D {
			return 0
		} else if d > D {
			return 1
		}
		return -1
	}

	w := 0
	for min <= max {
		w = (min + max) / 2
		r := can(w)
		if r == 0 {
			max = w - 1
		} else if r == 1 {
			min = w + 1
		} else {
			max = w - 1
		}
	}
	return min
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 3))
	fmt.Println(shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}
