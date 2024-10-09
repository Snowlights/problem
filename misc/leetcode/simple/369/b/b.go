package main

func minSum(nums1, nums2 []int) int64 {
	s1 := int64(0)
	zero1 := false
	for _, x := range nums1 {
		if x == 0 {
			zero1 = true
			s1++
		} else {
			s1 += int64(x)
		}
	}

	s2 := int64(0)
	zero2 := false
	for _, x := range nums2 {
		if x == 0 {
			zero2 = true
			s2++
		} else {
			s2 += int64(x)
		}
	}

	if zero1 && !zero2 && s1 > s2 ||
		!zero1 && zero2 && s1 < s2 ||
		!zero1 && !zero2 && s1 != s2 {
		return -1
	}
	return max(s1, s2)
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
