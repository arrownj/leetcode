package main

func Trap(height []int) int {
	dp := make([]int, len(height))
	maxBefore, minBefore := 0, 0
	for i, h := range height {
		if h <= height[minBefore] {
			dp[i] = dp[i-1]
			minBefore = i
		} else {
			if h >= height[maxBefore] {
				maxBefore = i
				minBefore = i
			}
			dp[i] = dp
		}

	}
	return 0
}
