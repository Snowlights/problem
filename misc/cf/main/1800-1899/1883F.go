//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1883F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		r := map[int]int{}
		for i := range a {
			Fscan(in, &a[i])
			r[a[i]] = i
		}

		ans := 0
		vis := map[int]bool{}
		rs := len(r)
		for i, v := range a {
			if !vis[v] {
				vis[v] = true
				ans += rs
			}
			if r[v] == i {
				rs--
			}
		}
		Fprintln(out, ans)
	}
}

func main() { CF1883F(os.Stdin, os.Stdout) }
