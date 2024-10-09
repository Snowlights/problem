package main

func countGoodNodes(edges [][]int) (ans int) {
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		size, pre, ok := 1, 0, true
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			sz := dfs(y, x)
			if pre > 0 && sz != pre {
				ok = false
			}
			pre = sz // 记录上一个儿子子树的大小
			size += sz
		}
		if ok {
			ans++
		}
		return size
	}
	dfs(0, -1)
	return
}
