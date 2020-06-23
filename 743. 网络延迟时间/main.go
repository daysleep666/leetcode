package main

import "fmt"

func networkDelayTime(times [][]int, N int, K int) int {
	const inifite = -1
	// 起始点到其他点的耗费时间
	node := make([]int, N+1)
	for i := 0; i <= N; i++ {
		node[i] = inifite
	}
	node[K] = 0

	// 任意一点到其他点 耗费的时间
	nodes := make([][][2]int, N+1)
	for _, v := range times {
		nodes[v[0]] = append(nodes[v[0]], [2]int{v[1], v[2]})
	}

	// 记录某点是否访问过
	visited := make(map[int]bool, N+1)

	// 从cur点开始找最短距离
	cur := K

	min := func(a, b int) int {
		if a == inifite || a > b {
			return b
		}
		return a
	}
	for cur != -1 { //
		visited[cur] = true
		// 计算最短路径
		for i := 0; i < len(nodes[cur]); i++ {
			other := nodes[cur][i][0]
			node[other] = min(node[other], node[cur]+nodes[cur][i][1])
		}
		next, min := -1, inifite
		// 选出下一个扫描的节点
		for i := 1; i < len(node); i++ {
			if visited[i] {
				continue
			}
			if node[i] != inifite && (min > node[i] || min == inifite) {
				next = i
				min = node[i]
			}
		}
		cur = next
	}

	max := 0
	for i := 1; i < len(node); i++ {
		if node[i] == inifite {
			return -1
		}
		if node[i] > max {
			max = node[i]
		}
	}
	return max
}

func main() {
	fmt.Println(networkDelayTime([][]int{[]int{
		2, 1, 1,
	}, []int{
		2, 3, 1,
	}, []int{
		3, 4, 1,
	}}, 4, 2))
}
