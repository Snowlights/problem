// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/diverta2019/submit?taskScreenName=diverta2019_c
// 对拍：https://atcoder.jp/contests/diverta2019/submissions?f.LanguageName=Go&f.Status=AC&f.Task=diverta2019_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3
			ABCA
			XBAZ
			BAD`,
			`2`,
		},
		{
			`9
			BEWPVCRWH
			ZZNQYIJX
			BAVREA
			PA
			HJMYITEOX
			BCJHMRMNK
			BP
			QVFABZ
			PRGKSPUNA`,
			`4`,
		},
		{
			`7
			RABYBBE
			JOZ
			BMHQUVA
			BPA
			ISU
			MCMABAOBHZ
			SZMEHMA`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/diverta2019/tasks/diverta2019_c
