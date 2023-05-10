package main

import "strings"

func IntToRoman(num int) string {
	if num <= 0 {
		panic("invalid input")
	}

	romanCharacterValues := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	romanSubtractCharacters := []string{"C", "X", "I"}
	romanCharacterLarger := map[string][]string{
		"C": []string{"M", "D"},
		"X": []string{"C", "L"},
		"I": []string{"X", "V"},
	}
	roman := ""
	for _, c := range romanSubtractCharacters {
		largerCharacters := romanCharacterLarger[c]
		if num >= romanCharacterValues[largerCharacters[0]] {
			roman += strings.Repeat(largerCharacters[0], num/romanCharacterValues[largerCharacters[0]])
			num = num % romanCharacterValues[largerCharacters[0]]
		}
		if num >= romanCharacterValues[c] {
			characterCount := num / romanCharacterValues[c]
			if characterCount == 9 {
				roman += c + largerCharacters[0]
			} else if characterCount >= 5 {
				roman += largerCharacters[1] + strings.Repeat(c, characterCount%5)
			} else if characterCount == 4 {
				roman += c + largerCharacters[1]
			} else {
				roman += strings.Repeat(c, characterCount%5)
			}
			num = num % romanCharacterValues[c]
		}
	}
	return roman
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
