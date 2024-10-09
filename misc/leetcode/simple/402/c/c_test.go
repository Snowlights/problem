// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_c(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, maximumTotalDamage, "c.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/weekly-contest-402/problems/maximum-total-damage-with-spell-casting/
// https://leetcode.cn/problems/maximum-total-damage-with-spell-casting/
