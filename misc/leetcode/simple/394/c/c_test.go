// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumOperations, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-394/problems/minimum-number-of-operations-to-satisfy-conditions/
// https://leetcode.cn/problems/minimum-number-of-operations-to-satisfy-conditions/
