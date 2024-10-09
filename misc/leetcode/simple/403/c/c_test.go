// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumTotalCost, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-403/problems/maximize-total-cost-of-alternating-subarrays/
// https://leetcode.cn/problems/maximize-total-cost-of-alternating-subarrays/
