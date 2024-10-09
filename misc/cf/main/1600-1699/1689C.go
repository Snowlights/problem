//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1689C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		g[1] = []int{0}
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			if len(g[v]) <= 2 {
				return len(g[v])
			}
			mx := int(1e9)
			for _, w := range g[v] {
				if w != fa {
					mx = min(mx, dfs(w, v)+2)
				}
			}
			return mx
		}
		Fprintln(out, n-dfs(1, 0))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1689C(os.Stdin, os.Stdout) }
