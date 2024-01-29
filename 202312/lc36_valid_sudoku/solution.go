package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool)
	columns := make(map[int]map[byte]bool)
	blocks := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			digit := board[i][j]
			if digit >= '1' && digit <= '9' {
				if row, ok := rows[i]; ok {
					if _, digitExist := row[digit]; digitExist {
						return false
					} else {
						row[digit] = true
					}
				} else {
					rows[i] = map[byte]bool{
						digit: true,
					}
				}

				if column, ok := columns[j]; ok {
					if _, digitExist := column[digit]; digitExist {
						return false
					} else {
						column[digit] = true
					}
				} else {
					columns[j] = map[byte]bool{
						digit: true,
					}
				}

				blockIndex := (i/3)*3 + j/3
				fmt.Printf("i: %d, j:%d, block: %d\n", i, j, blockIndex)
				if block, ok := blocks[blockIndex]; ok {
					fmt.Printf("block: %v\n", block)
					if _, digitExist := block[digit]; digitExist {
						return false
					} else {
						block[digit] = true
					}
				} else {
					blocks[blockIndex] = map[byte]bool{
						digit: true,
					}
				}
			}
		}
	}
	return true
}
