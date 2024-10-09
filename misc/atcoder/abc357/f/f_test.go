// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc357/submit?taskScreenName=abc357_f
// 对拍：https://atcoder.jp/contests/abc357/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc357_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`5 6
			1 3 5 6 8
			3 1 2 1 2
			3 1 3
			1 2 5 3
			3 1 3
			1 1 3 1
			2 5 5 2
			3 1 5`,
			`16
			25
			84`,
		},
		{
			`2 3
			1000000000 1000000000
			1000000000 1000000000
			3 1 1
			1 2 2 1000000000
			3 1 2`,
			`716070898
			151723988`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc357/tasks/abc357_f
