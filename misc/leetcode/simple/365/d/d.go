package main

func countVisitedNodes(g []int) []int {
	n := len(g)
	rg := make([][]int, n) // 反图
	deg := make([]int, n)
	for x, y := range g {
		rg[y] = append(rg[y], x)
		deg[y]++
	}

	// 拓扑排序，剪掉 g 上的所有树枝
	// 拓扑排序后，deg 值为 1 的点必定在基环上，为 0 的点必定在树枝上
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		y := g[x]
		deg[y]--
		if deg[y] == 0 {
			q = append(q, y)
		}
	}

	ans := make([]int, n)
	// 在反图上遍历树枝
	var rdfs func(int, int)
	rdfs = func(x, depth int) {
		ans[x] = depth
		for _, y := range rg[x] {
			if deg[y] == 0 { // 树枝上的点在拓扑排序后，入度均为 0
				rdfs(y, depth+1)
			}
		}
	}
	for i, d := range deg {
		if d <= 0 {
			continue
		}
		ring := []int{}
		for x := i; ; x = g[x] {
			deg[x] = -1            // 将基环上的点的入度标记为 -1，避免重复访问
			ring = append(ring, x) // 收集在基环上的点
			if g[x] == i {
				break
			}
		}
		for _, x := range ring {
			rdfs(x, len(ring)) // 为方便计算，以 len(ring) 作为初始深度
		}
	}
	return ans
}
