package main

func reverseWords(s string) string {
	words := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	if word != "" {
		words = append(words, word)
	}
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}
	return result
}
