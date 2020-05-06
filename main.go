package main

import (
	"math/rand"
	"sort"

	"github.com/daysleep666/leetcode/pkg"
)

func main() {
	arr := []int{4, 3, 2, 1}
	for i := 0; i < 10000; i++ {
		arr = append(arr, int(rand.Int63()))
	}
	pkg.Heapify(arr)
	// fmt.Println(arr)
	pkg.SortByHeap(arr)
	// fmt.Println(arr)
	if !sort.IntsAreSorted(arr) {
		panic("--")
	}
}
