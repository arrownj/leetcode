package main

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if length == 0 && s[i] == ' ' {
			continue
		} else if length > 0 && s[i] == ' ' {
			break
		}
		length++
	}
	return length
}
