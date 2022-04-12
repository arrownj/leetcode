package main

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	digitChars := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	combinations := []string{}
	firstDigit := string(digits[0])
	if letters, ok := digitChars[firstDigit]; ok {
		combinations = letters
	} else {
		panic("invalid input")
	}
	for i := 1; i < len(digits); i++ {
		digit := string(digits[i])
		if letters, ok := digitChars[digit]; ok {
			tmp := []string{}
			for j := 0; j < len(combinations); j++ {
				for k := 0; k < len(letters); k++ {
					tmp = append(tmp, combinations[j]+letters[k])
				}
			}
			combinations = tmp
		} else {
			panic("invalid input")
		}
	}
	return combinations
}
