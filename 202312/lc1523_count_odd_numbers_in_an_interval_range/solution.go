package main

func countOdds(low, high int) int {
	if low%2 == 0 {
		return (high + 1 - low) / 2
	}
	return (high-low)/2 + 1
}
