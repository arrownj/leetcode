package main

import "strings"

func intToRoman(num int) string {
	romanChars := []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	romanNumberMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	romanCharCount := make(map[byte]int)
	for _, c := range romanChars {
		romanCharCount[c] = num / romanNumberMap[c]
		num = num % romanNumberMap[c]
	}

	romanStringSlice := []byte{}
	for _, c := range romanChars {
		if count, ok := romanCharCount[c]; ok {
			for i := 0; i < count; i++ {
				romanStringSlice = append(romanStringSlice, c)
			}
		}
	}
	romanString := string(romanStringSlice)
	romanString = strings.ReplaceAll(romanString, "DCCCC", "CM")
	romanString = strings.ReplaceAll(romanString, "CCCC", "CD")
	romanString = strings.ReplaceAll(romanString, "LXXXX", "XC")
	romanString = strings.ReplaceAll(romanString, "XXXX", "XL")
	romanString = strings.ReplaceAll(romanString, "VIIII", "IX")
	romanString = strings.ReplaceAll(romanString, "IIII", "IV")
	return romanString
}
