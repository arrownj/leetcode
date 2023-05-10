package main

import "math"

func CountNumber(n int) int {
	count := 0
	for num := 0; num < int(math.Pow10(n)); num++ {
		if IsUniqueDigits(num) {
			count++
		}
	}
	return count
}

func IsUniqueDigits(num int) bool {
	remainders := make(map[int]bool)
	for num > 0 {
		remainder := num % 10
		if _, ok := remainders[remainder]; ok {
			return false
		} else {
			remainders[remainder] = true
		}
		num = num / 10
	}
	return true
}

func CountNumber2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	total, count := 10, 9
	for i := 1; i < n; i++ {
		count = count * (10 - i)
		total += count
	}

	return total
}
