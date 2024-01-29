package main

import "math"

func maxSubArray(nums []int) int {
	maxValues := make([]int, len(nums))
	maxValue := math.Inf(-1)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxValues[i] = nums[i]
		} else {
			maxValues[i] = max(nums[i], nums[i-1]+nums[i])
		}
		if maxValue < maxValues[i] {
			maxValue = maxValues[i]
		}
	}
	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
