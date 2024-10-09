// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc311/submit?taskScreenName=abc311_g
// 对拍：https://atcoder.jp/contests/abc311/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc311_g&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{
			`3 3
			5 4 3
			4 3 2
			3 2 1`,
			`48`,
		},
		{
			`4 5
			3 1 4 1 5
			9 2 6 5 3
			5 8 9 7 9
			3 2 3 8 4`,
			`231`,
		},
		{
			`6 6
			1 300 300 300 300 300
			300 1 300 300 300 300
			300 300 1 300 300 300
			300 300 300 1 300 300
			300 300 300 300 1 300
			300 300 300 300 300 1`,
			`810000`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc311/tasks/abc311_g
