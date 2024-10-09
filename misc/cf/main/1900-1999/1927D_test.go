//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1927/D
// https://codeforces.com/problemset/status/1927/problem/D
func TestCF1927D(t *testing.T) {
	t.Log("Current test is [CF1927D]")
	testCases := [][2]string{
		{
			`5
			5
			1 1 2 1 1
			3
			1 5
			1 2
			1 3
			6
			30 20 20 10 10 20
			5
			1 2
			2 3
			2 4
			2 6
			3 5
			4
			5 2 3 4
			4
			1 2
			1 4
			2 3
			2 4
			5
			1 4 3 2 4
			5
			1 5
			2 4
			3 4
			3 5
			4 5
			5
			2 3 1 4 2
			7
			1 2
			1 4
			1 5
			2 4
			2 5
			3 5
			4 5`,
			`2 3
			-1 -1
			1 3
			
			2 1
			-1 -1
			4 2
			4 6
			5 3
			
			1 2
			1 2
			2 3
			3 2
			
			1 3
			2 4
			3 4
			5 3
			5 4
			
			1 2
			4 2
			1 3
			2 3
			3 2
			5 4
			5 4
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1927D, testCases, 0)
}
