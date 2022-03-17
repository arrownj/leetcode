package main

func ClimbStairs(n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = 1
			continue
		}
		if i == 1 {
			dp[i] = 2
			continue
		}
		dp[i] = dp[i-1] + dp[i-2]

	}
	return dp[n-1]
}
