package main

func tictactoe(moves [][]int) string {
	grid := make([][]string, 3)
	for i := 0; i < 3; i++ {
		grid[i] = []string{"", "", ""}
	}
	for i := 0; i < len(moves); i++ {
		position := moves[i]
		row := position[0]
		col := position[1]
		owner := "A"
		if i%2 != 0 {
			owner = "B"
		}
		grid[row][col] = owner
		if hasWinner, winner := hasWinner(grid); hasWinner {
			return winner
		}
	}
	if len(moves) >= 9 {
		return "Draw"
	}
	return "Pending"
}

func hasWinner(grid [][]string) (bool, string) {
	for i := 0; i < 3; i++ {
		if (grid[i][0] == "A" && grid[i][1] == "A" && grid[i][2] == "A") || (grid[i][0] == "B" && grid[i][1] == "B" && grid[i][2] == "B") {
			return true, grid[i][0]
		}
		if (grid[0][i] == "A" && grid[1][i] == "A" && grid[2][i] == "A") || (grid[0][i] == "B" && grid[1][i] == "B" && grid[2][i] == "B") {
			return true, grid[0][i]
		}
	}
	if (grid[0][0] == "A" && grid[1][1] == "A" && grid[2][2] == "A") || (grid[0][0] == "B" && grid[1][1] == "B" && grid[2][2] == "B") || (grid[0][2] == "A" && grid[1][1] == "A" && grid[2][0] == "A") || (grid[0][2] == "B" && grid[1][1] == "B" && grid[2][0] == "B") {

		return true, grid[1][1]
	}
	return false, ""
}
