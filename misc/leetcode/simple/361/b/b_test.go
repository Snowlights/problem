// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, minimumOperations, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-361/problems/minimum-operations-to-make-a-special-number/
// https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/
