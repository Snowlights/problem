// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countGoodIntegers, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-138/problems/find-the-count-of-good-integers/
// https://leetcode.cn/problems/find-the-count-of-good-integers/
