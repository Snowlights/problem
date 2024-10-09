//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1209/F
// https://codeforces.com/problemset/status/1209/problem/F
func TestCF1209F(t *testing.T) {
	t.Log("Current test is [CF1209F]")
	testCases := [][2]string{
		{
			`11 10
			1 2
			2 3
			3 4
			4 5
			5 6
			6 7
			7 8
			8 9
			9 10
			10 11
			`,
			`1
			12
			123
			1234
			12345
			123456
			1234567
			12345678
			123456789
			345678826`,
		},
		{
			`12 19
			1 2
			2 3
			2 4
			2 5
			2 6
			2 7
			2 8
			2 9
			2 10
			3 11
			11 12
			1 3
			1 4
			1 5
			1 6
			1 7
			1 8
			1 9
			1 10
			`,
			`1
			12
			13
			14
			15
			16
			17
			18
			19
			1210
			121011
			`,
		},
		{
			`12 14
			1 2
			2 3
			3 4
			4 5
			5 6
			6 7
			7 8
			8 9
			9 10
			10 11
			11 12
			1 3
			1 4
			1 10
			`,
			`1
			12
			13
			134
			1345
			13456
			1498
			149
			14
			1410
			141011
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1209F, testCases, 0)
}
