// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_a(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, countCompleteDayPairs, "a.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-402/problems/count-pairs-that-form-a-complete-day-i/
// https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-i/