// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc281/submit?taskScreenName=abc281_f
// 对拍：https://atcoder.jp/contests/abc281/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc281_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`3
			12 18 11`,
			`16`,
		},
		{
			`10
			0 0 0 0 0 0 0 0 0 0`,
			`0`,
		},
		{
			`5
			324097321 555675086 304655177 991244276 9980291`,
			`805306368`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc281/tasks/abc281_f
