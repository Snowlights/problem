// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc294/submit?taskScreenName=abc294_E
// 对拍：https://atcoder.jp/contests/abc294/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc294_E&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [e]")
	testCases := [][2]string{
		{
			`8 4 3
			1 2
			3 2
			2 3
			3 1
			1 4
			2 1
			3 3`,
			`4`,
		},
		{
			`10000000000 1 1
			1 10000000000
			1 10000000000`,
			`10000000000`,
		},
		{
			`1000 4 7
			19 79
			33 463
			19 178
			33 280
			19 255
			33 92
			34 25
			19 96
			12 11
			19 490
			33 31`,
			`380`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc294/tasks/abc294_E
