//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

const mod = 998244353

func pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}

func CF1400G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	// 组合数模板
	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = F[i-1] * i % mod
	}
	invF := make([]int, n+1)
	invF[n] = pow(F[n], mod-2)
	for i := n; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
	C := func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	// 读入+计算差分
	a := make([]struct{ l, r int }, n+1)
	diff := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].l, &a[i].r)
		diff[a[i].l]++
		diff[a[i].r+1]--
	}
	ban := make([]struct{ x, y int }, m)
	for i := range ban {
		Fscan(in, &ban[i].x, &ban[i].y)
	}

	// 计算前缀和
	sum := make([][41]int, n+1)
	cnt := 0
	for i := 1; i <= n; i++ {
		cnt += diff[i]
		for j := 0; j < 41; j++ {
			sum[i][j] = (sum[i-1][j] + C(cnt-j, i-j)) % mod
		}
	}

	// 容斥
	ans := sum[n][0]
	has := make([]uint, n+1) // 避免使用哈希表的写法
	for i := uint(1); i < 1<<m; i++ {
		l, r, k := 1, n, 0
		for j := i; j > 0; j &= j - 1 {
			p := ban[bits.TrailingZeros(j)]
			// 计算区间交集，有多少个人
			l = max(l, a[p.x].l, a[p.y].l)
			r = min(r, a[p.x].r, a[p.y].r)
			if has[p.x] != i {
				has[p.x] = i // 时间戳
				k++
			}
			if has[p.y] != i {
				has[p.y] = i
				k++
			}
		}
		if r < l {
			continue
		}
		res := sum[r][k] - sum[l-1][k]
		if bits.OnesCount(i)%2 > 0 {
			res = -res
		}
		ans += res
	}
	Fprint(out, (ans%mod+mod)%mod) // 保证答案非负

}

func main() { CF1400G(os.Stdin, os.Stdout) }
