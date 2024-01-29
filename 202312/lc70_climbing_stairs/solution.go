package main

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ClimbStairs(n-2) + ClimbStairs(n-1)
}

func ClimbStairs(n int) int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ret[i] = 1
		} else if i == 1 {
			ret[i] = 2
		} else {
			ret[i] = ret[i-1] + ret[i-2]
		}
	}
	return ret[n-1]
}
