package main

func GenerateParentheses(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n <= 1 {
		return []string{"()"}
	}
	dp := [][]string{
		[]string{""},
		[]string{"()"},
	}
	for i := 2; i <= n; i++ {
		iResult := []string{}
		for p := 0; p < i; p++ {
			q := i - p - 1
			for _, pItem := range dp[p] {
				for _, qItem := range dp[q] {
					iResult = append(iResult, "("+pItem+")"+qItem)
				}
			}
		}
		dp = append(dp, iResult)
	}
	return dp[n]
}
