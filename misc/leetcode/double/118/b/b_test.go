// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximizeSquareHoleArea, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-118/problems/maximize-area-of-square-hole-in-grid/
// https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/
