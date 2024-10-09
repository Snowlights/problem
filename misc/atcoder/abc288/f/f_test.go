// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc288/submit?taskScreenName=abc288_f
// 对拍：https://atcoder.jp/contests/abc288/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc288_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`3
			234`,
			`418`,
		},
		{
			`4
			5915`,
			`17800`,
		},
		{
			`9
			998244353`,
			`258280134`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc288/tasks/abc288_f
