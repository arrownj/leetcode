package main

import (
	"math"
)

func numSquares(n int) int {
	nums := make([]int, n+1)
	nums[0] = 0
	for i := 1; i <= n; i++ {
		minNums := math.MaxInt
		for j := 1; j <= int(math.Sqrt(float64(i))); j++ {
			if minNums > 1+nums[i-j*j] {
				minNums = 1 + nums[i-j*j]
			}
		}
		nums[i] = minNums
	}
	return nums[n]
}
