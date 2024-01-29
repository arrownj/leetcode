package main

import "fmt"

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	results := make([][]int, len(num2))
	maxLength := 0
	for i := len(num2) - 1; i >= 0; i-- {
		result := []int{}
		for k := 0; k < len(num2)-1-i; k++ {
			result = append(result, 0)
		}
		carry := 0
		for j := len(num1) - 1; j >= 0; j-- {
			value := digit(num1[j])*digit(num2[i]) + carry
			result = append(result, value%10)
			carry = value / 10
		}
		if carry > 0 {
			result = append(result, carry)
		}
		if maxLength < len(result) {
			maxLength = len(result)
		}
		fmt.Printf("%v\n", result)
		results = append(results, result)
	}
	digits := []int{}
	carry := 0
	for i := 0; i < maxLength; i++ {
		sum := carry
		for _, result := range results {
			if i < len(result) {
				sum += result[i]
			}
		}
		carry = sum / 10
		digits = append(digits, sum%10)
	}
	if carry > 0 {
		digits = append(digits, carry)
	}
	result := ""
	for i := 0; i < len(digits); i++ {
		result = fmt.Sprintf("%v%s", digits[i], result)
	}
	return result
}

func digit(v byte) int {
	return int(v - '0')
}
