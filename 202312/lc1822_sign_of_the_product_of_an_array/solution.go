package main

func arraySign(nums []int) int {
	negativeCount := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negativeCount += 1
		}
	}
	if negativeCount%2 == 0 {
		return 1
	}
	return -1
}
