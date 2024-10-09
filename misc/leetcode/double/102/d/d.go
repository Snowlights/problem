package main

import "math"

const inf = math.MaxInt32 / 2 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（输入保证这条边之前不存在）
}

// 朴素 Dijkstra 算法
func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Floyd
//const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出
//
//type Graph [][]int
//
//func Constructor(n int, edges [][]int) Graph {
//	d := make([][]int, n) // 邻接矩阵
//	for i := range d {
//		d[i] = make([]int, n)
//		for j := range d[i] {
//			if j != i {
//				d[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
//			}
//		}
//	}
//	for _, e := range edges {
//		d[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边和自环）
//	}
//	for k := range d {
//		for i := range d {
//			for j := range d {
//				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
//			}
//		}
//	}
//	return d
//}
//
//func (d Graph) AddEdge(e []int) {
//	x, y, w := e[0], e[1], e[2]
//	if w >= d[x][y] { // 无需更新
//		return
//	}
//	for i := range d {
//		for j := range d {
//			d[i][j] = min(d[i][j], d[i][x]+w+d[y][j])
//		}
//	}
//}
//
//func (d Graph) ShortestPath(start, end int) int {
//	ans := d[start][end]
//	if ans == inf {
//		return -1
//	}
//	return ans
//}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
