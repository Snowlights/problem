//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1883/E
// https://codeforces.com/problemset/status/1883/problem/E
func TestCF1883E(t *testing.T) {
	t.Log("Current test is [CF1883E]")
	testCases := [][2]string{
		{
			`9
			1
			1
			2
			2 1
			3
			3 2 1
			4
			7 1 5 3
			5
			11 2 15 7 10
			6
			1 8 2 16 8 16
			2
			624323799 708290323
			12
			2 1 1 3 3 11 12 22 45 777 777 1500
			12
			12 11 10 9 8 7 6 5 4 3 2 1`,
			`0
			1
			3
			6
			10
			3
			0
			2
			66
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1883E, testCases, 0)
}
