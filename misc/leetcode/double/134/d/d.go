package main

import "sort"

func countSubarrays(nums []int, k int) (ans int64) {
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		a := nums[:i+1]
		ans += int64(sort.SearchInts(a, k+1) - sort.SearchInts(a, k))
	}
	return
}

func countSubarraysV2(nums []int, k int) (ans int64) {
	left, right := 0, 0
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		for left <= i && nums[left] < k {
			left++
		}
		for right <= i && nums[right] <= k {
			right++
		}
		ans += int64(right - left)
	}
	return
}
