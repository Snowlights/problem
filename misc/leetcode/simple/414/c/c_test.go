// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, findMaximumScore, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-414/problems/reach-end-of-array-with-max-score/
// https://leetcode.cn/problems/reach-end-of-array-with-max-score/
