package main

func CanJump(nums []int) bool {
	if len(nums) == 0 {
		panic("invalid input")
	}
	if len(nums) == 1 {
		return true
	}
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, len(nums))
		for j := 0; j < len(nums); j++ {
			if j > i && j <= i+nums[i] {
				dp[i][j] = true
				for k := 0; k < i; k++ {
					if dp[k][i] {
						dp[k][j] = true
					}
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[0][len(nums)-1]
}

func CanJump2(nums []int) bool {
	longest := 0
	for i := 0; i < len(nums) && i <= longest; i++ {
		if longest < i+nums[i] {
			longest = i + nums[i]
		}
		if longest >= len(nums)-1 {
			return true
		}
	}
	return false
}
