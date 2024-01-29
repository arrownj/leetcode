package main

import "math"

func MaxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if maxProfit < price-minPrice {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
