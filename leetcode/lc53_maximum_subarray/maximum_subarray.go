package main

func MaximumSubarray(nums []int) int {
	if len(nums) == 0 {
		panic("invalid input")
	}
	dp := make([]int, len(nums))
	max := nums[0]
	for i, num := range nums {
		if i == 0 {
			dp[i] = num
		} else {
			if dp[i-1] > 0 {
				dp[i] = dp[i-1] + num
			} else {
				dp[i] = num
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
