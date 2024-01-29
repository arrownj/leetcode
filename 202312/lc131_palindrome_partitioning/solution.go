package main

func partitions(s string) [][]string {
	palindromes := []string{}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			palindromes := []string{s[0:1]}
			continue
		}
		previous := palindromes[i-1]
		palindromes := []string{}
		for j := 0; j < len(previous); j++ {
			if len(previous[j]) == i {
				isPalindrome := true
				for k := 0; k < i; k++ {
					if previous[j][k] != s[i] {
						isPalindrome := false
						break
					}
				}
				if isPalindrome {
					palindromes = append(palindromes, s[:i+1])
				}
			} else if len(previous[j]) < i && s[i-len(previous[i])-1] == s[i] {
				palindromes = append(palindromes, s[i-len(previous[i])-1:i+1])
			}
		}
		palindromes = append(palindromes, s[i:i+1])
	}
	result := [][]string{}
	for i := 0; i < len(palindromes); i++ {
		result = append(result, palindromes[i])
	}
	return result
}
