package main

func canReachCorner(x, y int, a [][]int) bool {
	n := len(a)
	// 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
	fa := make([]int, n+2)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	merge := func(x, y int) {
		fa[find(x)] = find(y)
	}
	for i, c := range a {
		ox, oy, r := c[0], c[1], c[2]
		if ox <= r || oy+r >= y { // 圆 i 和左边界或上边界有交集
			merge(i, n)
		}
		if oy <= r || ox+r >= x { // 圆 i 和下边界或右边界有交集
			merge(i, n+1)
		}
		for j, q := range a[:i] {
			if (ox-q[0])*(ox-q[0])+(oy-q[1])*(oy-q[1]) <= (r+q[2])*(r+q[2]) {
				merge(i, j) // 圆 i 和圆 j 有交集
			}
		}
		if find(n) == find(n+1) { // 无法到达终点
			return false
		}
	}
	return true
}
