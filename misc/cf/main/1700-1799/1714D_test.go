//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1714/D
// https://codeforces.com/problemset/status/1714/problem/D
func TestCF1714D(t *testing.T) {
	t.Log("Current test is [CF1714D]")
	testCases := [][2]string{
		{
			`6
			bababa
			2
			ba
			aba
			caba
			2
			bac
			acab
			abacabaca
			3
			aba
			bac
			aca
			baca
			3
			a
			c
			b
			codeforces
			4
			def
			code
			efo
			forces
			aaaabbbbcccceeee
			4
			eeee
			cccc
			aaaa
			bbbb`,
			`3
			2 2
			1 1
			2 4
			-1
			4
			1 1
			2 6
			3 3
			3 7
			4
			3 1
			1 2
			2 3
			1 4
			2
			4 5
			2 1
			4
			3 1
			4 5
			2 9
			1 13`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1714D, testCases, 0)
}
