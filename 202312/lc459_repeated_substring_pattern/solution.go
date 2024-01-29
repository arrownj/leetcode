package main

func isRepeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			isRepeated := true
			j := i
			for j < len(s) {
				if s[j:j+i] != s[0:i] {
					isRepeated = false
					break
				}
				j += i
			}
			if isRepeated {
				return true
			}
		}
	}
	return false
}
