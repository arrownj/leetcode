package main

import (
	"math"
)

func RobRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	return int(math.Max(float64(Rob(nums[:len(nums)-1])), float64(Rob(nums[:len(nums)-2])+nums[len(nums)-1])))
}

func RobIterative(nums []int) int {
	maxRob := []int{}
	for idx, num := range nums {
		if idx == 0 {
			maxRob = append(maxRob, num)
		} else if idx == 1 {
			maxRob = append(maxRob, int(math.Max(float64(nums[0]), float64(nums[1]))))
		} else {
			maxRob = append(maxRob, int(math.Max(float64(maxRob[idx-1]), float64(maxRob[idx-2]+num))))
		}
	}
	return maxRob[len(maxRob)-1]
}
