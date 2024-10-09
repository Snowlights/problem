package main

func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	ans := cnt[1] - (cnt[1]%2 ^ 1)
	delete(cnt, 1)
	for x := range cnt {
		res := 0
		for ; cnt[x] > 1; x *= x {
			res += 2
		}
		res += cnt[x]
		ans = max(ans, res-(res%2^1)) // 保证 res 是奇数
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
