package main

import "math"

func averageSalary(salary []int) float64 {
	minSalary, maxSalary := math.MaxInt, math.MinInt
	sum := 0
	for i := 0; i < len(salary); i++ {
		if minSalary > salary[i] {
			minSalary = salary[i]
		}
		if maxSalary < salary[i] {
			maxSalary = salary[i]
		}
		sum += salary[i]
	}
	return float64(sum-minSalary-maxSalary) / float64(len(salary)-2)
}
