//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1515/F
// https://codeforces.com/problemset/status/1515/problem/F
func TestCF1515F(t *testing.T) {
	t.Log("Current test is [CF1515F]")
	testCases := [][2]string{
		{
			`5 4 1
			0 0 0 4 0
			1 2
			2 3
			3 4
			4 5
			`,
			`YES
			3
			2
			1
			4
			`,
		},
		{
			`2 1 2
			1 1
			1 2
			`,
			`YES
			1
			`,
		},
		{
			`2 1 2
			0 1
			1 2
			`,
			`NO`,
		},
		{
			`5 6 5
			0 9 4 0 10
			1 2
			1 3
			2 3
			3 4
			1 4
			4 5
			`,
			`YES
			6
			4
			1
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1515F, testCases, 0)
}
