package main

import "math"

func maximumWealth(accounts [][]int) int {
	maximum := math.MinInt
	for i := 0; i < len(accounts); i++ {
		wealth := 0
		for j := 0; j < len(accounts[i]); j++ {
			wealth += accounts[i][j]
		}
		if maximum < wealth {
			maximum = wealth
		}
	}
	return maximum
}
