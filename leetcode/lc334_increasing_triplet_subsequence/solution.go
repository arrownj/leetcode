package lc334_increasing_triplet_subsequence

import "math"

func Solution(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first := nums[0]
	second := math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i] > second {
			return true
		} else if nums[i] > first {
			second = numns[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
