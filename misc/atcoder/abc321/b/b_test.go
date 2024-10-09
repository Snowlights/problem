// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc321/submit?taskScreenName=abc321_b
// 对拍：https://atcoder.jp/contests/abc321/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc321_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`5 180
			40 60 80 50`,
			`70`,
		},
		{
			`3 100
			100 100`,
			`0`,
		},
		{
			`5 200
			0 0 99 99`,
			`-1`,
		},
		{
			`10 480
			59 98 88 54 70 24 8 94 46`,
			`45`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc321/tasks/abc321_b
