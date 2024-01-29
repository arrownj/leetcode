package main

func setZeroes(matrix [][]int) {
	if len(matrix) > 0 {
		rows := make(map[int]bool)
		columms := make(map[int]bool)
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] == 0 {
					rows[i] = true
					columns[j] = true
				}
			}
		}
	}
}
