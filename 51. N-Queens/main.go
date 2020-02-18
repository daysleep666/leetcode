// package main

// import "fmt"

// func solveNQueens(n int) (result [][]string) {
// 	type point struct {
// 		x, y int
// 	}
// 	var ps []*point
// 	check := func(x, y int) bool {
// 		for _, p := range ps {
// 			if p.x == x || p.y == y {
// 				return false
// 			}
// 			if (p.x > x && p.y > y) || (p.x < x && p.y < y) {
// 				if (p.y - p.x) == (y - x) {
// 					return false
// 				}
// 			} else {
// 				if (p.x + p.y) == (x + y) {
// 					return false
// 				}
// 			}
// 		}
// 		return true
// 	}

// 	// for k := 0; (n%2 == 0 && k < n/2) || (n%2 != 0 && k <= n/2); k++ {
// 	for k := 0; k < n; k++ {
// 		r := make([]string, n)
// 		ps = append(ps, &point{x: 0, y: k})
// 		result[0][k] = "Q"

// 		i := 1
// 		for ; i < n; i++ {
// 			var has bool
// 			for j := 0; j < n; j++ {
// 				if check(i, j) {
// 					ps = append(ps, &point{x: i, y: j})
// 					result[i][j] = "Q"
// 					has = true
// 					break
// 				} else {
// 					result[i][j] = "."
// 				}
// 			}
// 			if !has {
// 				break
// 			}
// 		}
// 		if i == n {
// 			return
// 		}
// 	}
// 	return
// }

// func main() {
// 	print := func(n int) {
// 		r := solveNQueens(n)
// 		for i := 0; i < n; i++ {
// 			for j := 0; j < n; j++ {
// 				fmt.Printf("%s ", r[i][j])
// 			}
// 			fmt.Println()
// 		}
// 	}
// 	print(4)
// }
