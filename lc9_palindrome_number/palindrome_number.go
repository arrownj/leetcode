package main

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	value, inputX := 0, x
	for x > 0 {
		value *= 10
		value += x % 10
		x /= 10
	}
	return inputX == value
}
