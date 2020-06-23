package main

import (
	"fmt"
)

// 深度优先搜索（DFS）+ 剪枝
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if recursive(board, word, i, j) {
				return true
			}
		}
	}
	return false
}

// 从指定点出发,寻找对应的子字符串
// 深度优先遍历,走完一条路后要复原整条路
func recursive(board [][]byte, word string, i, j int) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}
	cur := word[0]
	// 将走过的点标记下
	board[i][j] = '/'
	defer func() {
		board[i][j] = cur
	}()
	word = word[1:]
	return recursive(board, word, i+1, j) ||
		recursive(board, word, i-1, j) ||
		recursive(board, word, i, j+1) ||
		recursive(board, word, i, j-1)
}

func main() {
	fmt.Println(exist([][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'F', 'S', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	fmt.Println(exist([][]byte{
		[]byte{'a', 'b'},
	}, "ba"))
}
