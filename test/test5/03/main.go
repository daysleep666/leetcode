package main

// dijs
// 到达某一点的距离是最短的
func networkDelayTime(times [][]int, N int, K int) int {
	// 节点到节点的权重
	graph := make(map[[2]int]int, 0)

	// 访问过的节点
	visited := make(map[int]int, N)

	// 节点的边关系
	node := make(map[int][][2]int, N)
	for _, v := range times {
		node[v] = append(node[v[0]], [2]int{v[1], v[2]})
	}

	// 当前访问的节点
	curNode := k
	for {
		// cur节点到其他节点的权重
		n := node[curNode]
		for i := 0; i < len(n); i++ {
			other, weight := n[0], n[1]
			// 更新当前节点到其他节点的权重
			graph[[2]int{curNode, other}] = weight
			// 更新初始节点到当前节点再到其他节点的权重
			// graph[[2]int{K,}]
		}

	}
}

func main() {
	// fmt.Println(networkDelayTime())
}
