package _250

import (
	"bufio"
	"fmt"
	"io"
)

// https://codeforces.com/problemset/problem/1286/A
//
//输入 n(≤100) 和一个长为 n 的数组 p，p 原本是一个 1~n 的排列，但是有些数字丢失了，丢失的数字用 0 表示。
//你需要还原 p，使得 p 中相邻元素奇偶性不同的对数最少。输出这个最小值。

// 定义 f[i][j][0/1] 表示前 i 个数中有 j 个偶数时的奇偶性不同对的最小个数，
// 其中 f[i][j][0] 表示第 i 个数是偶数，f[i][j][1] 表示第 i 个数是奇数。
//讨论当前数字和上一个数字的奇偶性，得到转移方程（如果能填奇数/偶数才能转移）：
// f[i+1][j][0] = min(f[i][j−1][0],f[i][j−1][1]+1)
// f[i+1][j][1] = min(f[i][j][0]+1, f[i][j][1])
// 初始项 f[0][0][0]=0,f[0][0][1]=0，其余为无穷大。
//答案为 min(f[n][n/2][0],f[n][n/2][1])

func CF1286A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, v int
	fmt.Fscan(r, &n)

	f := make([][][2]int, n+1)
	for i := range f {
		f[i] = make([][2]int, n/2+1)
		for j := range f[i] {
			f[i][j] = [2]int{1e9, 1e9}
		}
	}
	f[0][0] = [2]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &v)
		for j := 0; j <= n/2; j++ {
			if j > 0 && v%2 == 0 {
				f[i+1][j][0] = min(f[i][j-1][0], f[i][j-1][1]+1)
			}
			if v == 0 || v%2 > 0 {
				f[i+1][j][1] = min(f[i][j][0]+1, f[i][j][1])
			}
		}
	}
	fmt.Fprintln(w, min(f[n][n/2][0], f[n][n/2][1]))

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
