// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc243/submit?taskScreenName=abc243_g
// 对拍：https://atcoder.jp/contests/abc243/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc243_g&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [g]")
	testCases := [][2]string{
		{
			`4
			16
			1
			123456789012
			1000000000000000000`,
			`5
			1
			4555793983
			23561347048791096`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc243/tasks/abc243_g