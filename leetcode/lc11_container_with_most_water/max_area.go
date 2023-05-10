package main

func MaxArea(height []int) int {
	if len(height) < 2 {
		panic("invalid input")
	}
	max := 0
	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			volume := (i - j) * min(height[i], height[j])
			if max < volume {
				max = volume
			}
		}
	}
	return max
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}

func MaxArea2(height []int) int {
	if len(height) < 2 {
		panic("invalid input")
	}
	max := 0
	for i := 1; i < len(height); i++ {
		for j := 0; j < i; j++ {
			volume := (i - j) * min(height[i], height[j])
			if max < volume {
				max = volume
			}
		}
	}
	return max
}
