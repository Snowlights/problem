// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, leftmostBuildingQueriesSeg, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-372/problems/find-building-where-alice-and-bob-can-meet/
// https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet/
