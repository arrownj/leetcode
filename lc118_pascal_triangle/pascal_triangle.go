package main

func PascalTriangle(numRow int) [][]int {
	if numRow <= 0 {
		panic("invalid input")
	}
	rows := make([][]int, numRow)
	for i := 1; i <= numRow; i++ {
		if i == 1 {
			rows[i-1] = []int{1}
			continue
		}
		row := make([]int, i)
		for j := 0; j < i; j++ {
			leftIndex := j - 1
			rightIndex := j
			if leftIndex < 0 {
				row[j] = rows[i-2][rightIndex]
			} else if rightIndex > i-2 {
				row[j] = rows[i-2][leftIndex]
			} else {
				row[j] = rows[i-2][leftIndex] + rows[i-2][rightIndex]
			}
		}
		rows[i-1] = row
	}
	return rows
}
