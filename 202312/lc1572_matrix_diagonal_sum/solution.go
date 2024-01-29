package main

func diagonalSum(mat [][]int) int {
	length := len(mat)
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		sum += mat[i][length-i-1]
	}
	if length%2 != 0 {
		sum -= mat[(length-1)/2][(length-1)/2]
	}
	return sum
}
