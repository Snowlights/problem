//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF797F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, sumCap int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	slices.Sort(a)
	b := make([]struct{ x, cap int }, m)
	for i := range b {
		Fscan(in, &b[i].x, &b[i].cap)
		sumCap += b[i].cap
	}
	if sumCap < n {
		Fprint(out, -1)
		return
	}
	sort.Slice(b, func(i, j int) bool { return b[i].x < b[j].x })

	f := make([]int, n+1)
	for j := 1; j <= n; j++ {
		f[j] = 1e18
	}
	s := make([]int, n+1)
	for _, p := range b {
		for j, x := range a {
			s[j+1] = s[j] + abs(x-p.x)
		}
		type pair struct{ v, i int }
		q := []pair{{}}
		for j := 1; j <= n; j++ {
			for len(q) > 0 && f[j]-s[j] <= q[0].v {
				q = q[:len(q)-1]
			}
			q = append(q, pair{f[j] - s[j], j})
			if q[0].i < j-p.cap {
				q = q[1:]
			}
			f[j] = q[0].v + s[j]
		}
	}
	Fprint(out, f[n])

}

func abs(x int) int { if x < 0 { return -x }; return x }


func main() { CF797F(os.Stdin, os.Stdout) }
