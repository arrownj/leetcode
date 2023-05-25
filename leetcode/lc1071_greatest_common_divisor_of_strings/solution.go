package lc1071_greatest_common_divisor_of_strings

import (
	"math"
)

func Solution(s1, s2 string) string {
	for i := int(math.Min(float64(len(s1)), float64(len(s2)))); i > 0; i-- {
		if len(s1)%i == 0 && len(s2)%i == 0 {
			if s1 == multiply(s1[:i], len(s1)/i) && s2 == multiply(s1[:i], len(s2)/i) {
				return s1[:i]
			}
		}
	}
	return ""
}

func multiply(s string, num int) string {
	output := ""
	for i := 0; i < num; i++ {
		output += s
	}
	return output
}
