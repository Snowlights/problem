package main

func duplicateNumbersXOR(nums []int) (ans int) {
	vis := 0
	for _, x := range nums {
		if vis>>x&1 > 0 {
			ans ^= x
		} else {
			vis |= 1 << x
		}
	}
	return
}
