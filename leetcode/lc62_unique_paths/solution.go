package lc62_unique_paths

func Solution(m, n int) int {
	if m <= 0 || n <= 0 {
		panic("invalid input")
	}
	paths := make([][]int, m, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				paths[i][j] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[m-1][n-1]
}
