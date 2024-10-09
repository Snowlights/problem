// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc159/submit?taskScreenName=abc159_b
// 对拍：https://atcoder.jp/contests/abc159/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc159_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`akasaka`,
			`Yes`,
		},
		{
			`level`,
			`No`,
		},
		{
			`atcoder`,
			`No`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc159/tasks/abc159_b
