package main

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	if x-1 >= 0 && grid[x-1][y] == '1' {
		dfs(grid, x-1, y)
	}
	if x+1 < len(grid) && grid[x+1][y] == '1' {
		dfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		dfs(grid, x, y-1)
	}
	if y+1 < len(grid[0]) && grid[x][y+1] == '1' {
		dfs(grid, x, y+1)
	}
}
