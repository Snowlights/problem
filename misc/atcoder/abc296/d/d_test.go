// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc296/submit?taskScreenName=abc296_d
// 对拍：https://atcoder.jp/contests/abc296/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc296_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`5 7`,
			`8`,
		},
		{
			`2 5`,
			`-1`,
		},
		{
			`100000 10000000000`,
			`10000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc296/tasks/abc296_d
