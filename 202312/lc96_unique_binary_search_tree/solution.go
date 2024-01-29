package main

func numTrees(n int) int {
	nums := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 || i == 1 {
			nums[i] = 1
			continue
		}
		num := 0
		for j := 0; j < i; j++ {
			num += nums[j] * nums[i-j-1]
		}
		nums[i] = num
	}
	return nums[n]
}
