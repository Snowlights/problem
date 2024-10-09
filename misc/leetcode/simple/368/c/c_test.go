// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minGroupsForValidAssignment, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-368/problems/minimum-number-of-groups-to-create-a-valid-assignment/
// https://leetcode.cn/problems/minimum-number-of-groups-to-create-a-valid-assignment/