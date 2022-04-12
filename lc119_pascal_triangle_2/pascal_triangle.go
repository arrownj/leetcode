package main

func PascalTriangle(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		last := result[0]
		for j := 1; j <= i; j++ {
			tmp := result[j]
			if j != i {
				result[j] = last + result[j]
			} else {
				result[j] = 1
			}
			last = tmp
		}
	}
	return result
}
