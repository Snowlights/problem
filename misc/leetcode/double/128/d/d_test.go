// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	if err := leetcode.RunLeetCodeFuncWithFile(t, numberOfSubarrays, "d.txt", 0); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-128/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/
// https://leetcode.cn/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/
