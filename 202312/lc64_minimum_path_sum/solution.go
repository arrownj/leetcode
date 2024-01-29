package main

func minPathSum(grid [][]int) int {
	sums := make([][]int, len(grid))
	for i := 0; i < len(sums); i++ {
		sums[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				if j == 0 {
					sums[i][j] = grid[i][j]
				} else {
					sums[i][j] = sums[i][j-1] + grid[i][j]
				}
			} else if j == 0 {
				sums[i][j] = sums[i-1][j] + grid[i][j]
			} else {
				sums[i][j] = min(sums[i-1][j], sums[i][j-1]) + grid[i][j]
			}
		}
	}
	return sums[len(grid)-1][len(grid[0])-1]
}
