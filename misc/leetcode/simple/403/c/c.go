package main

import (
	"math"
)

func maximumTotalCost(a []int) int64 {
	f0, f1 := 0, 0
	for i := len(a) - 1; i >= 0; i-- {
		f0, f1 = f1+a[i], max(f1+a[i], f0-a[i])
	}
	return int64(f0)
}

func maximumTotalCost1(a []int) int64 {
	n := len(a)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{math.MinInt, math.MinInt}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == n {
			return 0
		}
		p := &memo[i][j]
		if *p != math.MinInt {
			return *p
		}
		res := dfs(i+1, 1) + a[i]
		r := dfs(i+1, j^1)
		if j == 0 {
			r += a[i]
		} else {
			r -= a[i]
		}
		res = max(res, r)
		*p = res
		return res
	}
	return int64(dfs(0, 0))
}
