package main

func MinimumTotal(triangle [][]int) int {
	minRows := make([][]int, len(triangle))
	for i := 0; i < len(minRows); i++ {
		if i == 0 {
			minRows[i] = []int{triangle[0][0]}
			continue
		}
		minRow := make([]int, i+1)
		for j := 0; j < len(minRow); j++ {
			if j == 0 {
				minRow[j] = triangle[i][j] + minRows[i-1][j]
			} else if j == len(minRow)-1 {
				minRow[j] = triangle[i][j] + minRows[i-1][j-1]
			} else {
				if minRows[i-1][j-1] < minRows[i-1][j] {
					minRow[j] = triangle[i][j] + minRows[i-1][j-1]
				} else {
					minRow[j] = triangle[i][j] + minRows[i-1][j]
				}
			}
		}
		minRows[i] = minRow
	}
	minTotal := minRows[len(minRows)-1][0]
	for k := 0; k < len(minRows[len(minRows)-1]); k++ {
		if minRows[len(minRows)-1][k] < minTotal {
			minTotal = minRows[len(minRows)-1][k]
		}
	}
	return minTotal
}
