package main

import (
	"math"
	"strings"
)

func Atoi(s string) int {
	isPositive := true
	MIN, MAX := -int(math.Pow(2, 31)), int(math.Pow(2, 31)-1)
	value := 0
	s = strings.Trim(s, " ")
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (i > 0 && !IsDigit(c)) || (i == 0 && !IsDigit(c) && !IsSign(c)) {
			break
		}
		if i == 0 && c == '-' {
			isPositive = false
		}
		if IsDigit(c) {
			value *= 10
			value += ByteToInt(c)
		}
		if isPositive && value > MAX {
			return MAX
		} else if !isPositive && value > MAX+1 {
			return MIN
		}
	}
	if !isPositive {
		value = -value
	}
	return value
}

func IsDigit(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' ||
		c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

func IsSign(c byte) bool {
	return c == '-' || c == '+'
}

func ByteToInt(c byte) int {
	if c == '0' {
		return 0
	}
	if c == '1' {
		return 1
	}
	if c == '2' {
		return 2
	}
	if c == '3' {
		return 3
	}
	if c == '4' {
		return 4
	}
	if c == '5' {
		return 5
	}
	if c == '6' {
		return 6
	}
	if c == '7' {
		return 7
	}
	if c == '8' {
		return 8
	}
	if c == '9' {
		return 9
	}
	return 0
}
