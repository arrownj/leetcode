package divide_two_integers

import "math"

func Divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("can't divide 0")
	}
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	sign := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}
	quotient := 0
	for divisor <= dividend {
		index := 1
		sub_divisor := divisor
		for sub_divisor <= dividend {
			dividend -= sub_divisor
			sub_divisor += sub_divisor
			quotient += index
			index += index
		}
	}
	if sign == -1 {
		quotient = 0 - quotient
	}
	return quotient
}
