// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxIncreasingCells, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-347/problems/maximum-strictly-increasing-cells-in-a-matrix/
// https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/