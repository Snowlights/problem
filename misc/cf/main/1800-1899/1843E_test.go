//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1843/E
// https://codeforces.com/problemset/status/1843/problem/E
func TestCF1843E(t *testing.T) {
	t.Log("Current test is [CF1843E]")
	testCases := [][2]string{
		{
			`6
			5 5
			1 2
			4 5
			1 5
			1 3
			2 4
			5
			5
			3
			1
			2
			4
			4 2
			1 1
			4 4
			2
			2
			3
			5 2
			1 5
			1 5
			4
			2
			1
			3
			4
			5 2
			1 5
			1 3
			5
			4
			1
			2
			3
			5
			5 5
			1 5
			1 5
			1 5
			1 5
			1 4
			3
			1
			4
			3
			3 2
			2 2
			1 3
			3
			2
			3
			1`,
			`3
			-1
			3
			3
			3
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1843E, testCases, 0)
}
