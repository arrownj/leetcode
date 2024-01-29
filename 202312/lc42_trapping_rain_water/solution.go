package main

import (
	"fmt"
	"math"
)

func trap(height []int) int {
	fmt.Printf("%v\n", height)
	totalWater := 0
	for idx := 0; idx < len(height); idx++ {
		maxLeft, maxRight := 0, 0
		for i := idx - 1; i >= 0; i-- {
			if maxLeft < height[i] {
				maxLeft = height[i]
			}
		}
		for j := idx + 1; j < len(height); j++ {
			if maxRight < height[j] {
				maxRight = height[j]
			}
		}
		fmt.Printf("%d, %d\n", maxLeft, maxRight)
		minMax := int(math.Min(float64(maxLeft), float64(maxRight)))
		if height[idx] < minMax {
			totalWater += minMax - height[idx]
		}
	}
	return totalWater
}
