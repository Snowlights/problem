// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maxTotalReward, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-401/problems/maximum-total-reward-using-operations-i/
// https://leetcode.cn/problems/maximum-total-reward-using-operations-i/