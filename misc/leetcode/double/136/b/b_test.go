// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minFlips, "b.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-136/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/
// https://leetcode.cn/problems/minimum-number-of-flips-to-make-binary-grid-palindromic-i/
