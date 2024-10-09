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

type data struct{ c0, c1, c01, c10 int }

type seg []struct {
	l, r int
	d    data
	flip bool
}

func mergeInfo(a, b data) (c data) {
	c.c0 = a.c0 + b.c0
	c.c1 = a.c1 + b.c1
	c.c01 = max(a.c01+b.c1, a.c0+b.c01)
	c.c10 = max(a.c10+b.c0, a.c1+b.c10)
	return
}

func (t seg) do(O int) {
	o := &t[O]
	d := o.d
	o.d = data{d.c1, d.c0, d.c10, d.c01}
	o.flip = !o.flip
}

func (t seg) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		c1 := int(s[l-1] & 1)
		t[o].d = data{c1 ^ 1, c1, 1, 1}
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	if t[o].flip {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func CF145E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	var s string
	Fscan(in, &n, &m, &s)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(s, 1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &s)
		if s[0] == 's' {
			Fscan(in, &l, &r)
			t.update(1, l, r)
		} else {
			Fprintln(out, t[1].d.c01)
		}
	}
	
}

func main() { CF145E(os.Stdin, os.Stdout) }
