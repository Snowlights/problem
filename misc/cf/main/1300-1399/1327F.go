//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1327F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mod = 998244353
	var n, k, m int
	Fscan(in, &n, &k, &m)
	a := make([]struct{ l, r, x int }, m)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r, &a[i].x)
	}
	ans := 1
	for b := 0; b < k; b++ {
		maxL := make([]int, n+1)
		d := make([]int, n+2)
		for _, p := range a {
			if p.x>>b&1 == 0 {
				maxL[p.r] = max(maxL[p.r], p.l)
			} else {
				d[p.l]++
				d[p.r+1]--
			}
		}

		f := make([]int, n+2)
		f[0] = 1
		sf := 1
		sd := 0
		l := 0
		for i := 1; i <= n+1; i++ {
			for l < maxL[i-1] {
				sf -= f[l]
				l++
			}
			sd += d[i]
			if sd > 0 {
				continue
			}
			sf %= mod
			f[i] = sf
			sf = sf * 2
		}
		ans = ans * f[n+1] % mod
	}
	ans = (ans%mod + mod) % mod
	Fprint(out, ans)

}

func main() { CF1327F(os.Stdin, os.Stdout) }
