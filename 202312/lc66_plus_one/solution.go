package main

func plusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if i == len(digits)-1 {
			sum = sum + 1
		}
		carry = sum / 10
		result[i+1] = sum % 10
	}
	if carry > 0 {
		result[0] = carry
	} else {
		result = result[1:]
	}
	return result
}
