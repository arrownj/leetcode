package main

func generateParenthesis(n int) []string {
	parenthesises := make([][]string, n+1)
	parenthesises[0] = []string{}
	for i := 1; i <= n; i++ {
		curr := []string{}
		for j := 0; j < i; j++ {
			left := parenthesises[j]
			right := parenthesises[i-j-1]
			for _, l := range left {
				for _, r := range right {
					curr = append(curr, "("+l+")"+r)
				}
			}
		}
		parenthesises[i] = curr
	}
	return parenthesises[n]
}
