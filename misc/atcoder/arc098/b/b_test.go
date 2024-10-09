// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc098/submit?taskScreenName=arc098_b
// 对拍：https://atcoder.jp/contests/arc098/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc098_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`4
			2 5 4 6`,
			`5`,
		},
		{
			`9
			0 0 0 0 0 0 0 0 0`,
			`45`,
		},
		{
			`19
			885 8 1 128 83 32 256 206 639 16 4 128 689 32 8 64 885 969 1`,
			`37`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc098/tasks/arc098_b