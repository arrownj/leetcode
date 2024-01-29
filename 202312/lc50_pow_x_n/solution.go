package main

import "math"

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0.0
	}
	ret := x
	i := 1
	for i <= int(math.Abs(float64(n))) {
		ret *= ret
		i = i * 2
	}
	if n < 0 {
		ret = 1.0 / ret
	}
	return ret
}
