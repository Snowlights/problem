// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc288/submit?taskScreenName=abc288_g
// 对拍：https://atcoder.jp/contests/abc288/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc288_g&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{
			`1
			0 1 1`,
			`0 0 1`,
		},
		{
			`2
			2 3 2 4 5 3 3 4 2`,
			`0 1 0 1 0 1 1 1 0`,
		},
		{
			`2
			0 0 0 0 0 0 0 0 0`,
			`0 0 0 0 0 0 0 0 0`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc288/tasks/abc288_g
