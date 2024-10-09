// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc159/submit?taskScreenName=abc159_a
// 对拍：https://atcoder.jp/contests/abc159/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc159_a&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`2 1`,
			`1`,
		},
		{
			`4 3`,
			`9`,
		},
		{
			`1 1`,
			`0`,
		},
		{
			`13 3`,
			`81`,
		},
		{
			`0 3`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc159/tasks/abc159_a