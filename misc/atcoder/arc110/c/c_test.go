// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc110/submit?taskScreenName=arc110_c
// 对拍：https://atcoder.jp/contests/arc110/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc110_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`5
			2 4 1 5 3`,
			`4
			2
			3
			1`,
		},
		{
			`5
			5 4 3 2 1`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/arc110/tasks/arc110_c
