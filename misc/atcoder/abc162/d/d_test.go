// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc162/submit?taskScreenName=abc162_d
// 对拍：https://atcoder.jp/contests/abc162/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc162_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`4
			RRGB`,
			`1`,
		},
		{
			`39
			RBRBGRBGGBBRRGBBRRRBGGBRBGBRBGBRBBBGBBB`,
			`1800`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc162/tasks/abc162_d