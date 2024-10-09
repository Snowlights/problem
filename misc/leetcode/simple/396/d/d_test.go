package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minCostToEqualizeArray, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-396/problems/minimum-cost-to-equalize-array/
// https://leetcode.cn/problems/minimum-cost-to-equalize-array/
