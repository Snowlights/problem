// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minAbsoluteDifference, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-358/problems/minimum-absolute-difference-between-elements-with-constraint/
// https://leetcode.cn/problems/minimum-absolute-difference-between-elements-with-constraint/
