package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	type s struct {
		n int
		c int
	}
	ss := make([]*s, 0, len(nums))
	add := func(n int) {
		i := 0
		for ; i < len(ss); i++ {
			if ss[i].n == n {
				ss[i].c++
				break
			}
		}
		if i == len(ss) {
			ss = append(ss, &s{n: n, c: 1})
			return
		}
		for ; i > 0; i = (i - 1) / 2 {
			if ss[i].c > ss[(i-1)/2].c {
				ss[i], ss[(i-1)/2] = ss[(i-1)/2], ss[i]
			}
		}
	}
	pop := func() int {
		r := ss[0]
		for i := 0; i*2+1 < len(ss); {
			if i*2+2 >= len(ss) || ss[i*2+1].c > ss[i*2+2].c {
				ss[i], ss[i*2+1] = ss[i*2+1], ss[i]
				i = i*2 + 1
			} else if i*2+2 < len(ss) {
				ss[i], ss[i*2+2] = ss[i*2+2], ss[i]
				i = i*2 + 2
			}
		}
		return r.n
	}
	for _, n := range nums {
		add(n)
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pop()
	}
	return result
}

func main() {
	fmt.Println(topKFrequent([]int{5, 2, 5, 3, 5, 3, 1, 1, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}, 3))
}
