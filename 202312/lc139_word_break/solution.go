package main

func wordBreak(s string, wordDict []string) bool {
	canBreaks := make([]bool, len(s)+1)
	canBreaks[0] = true
	for i := 1; i <= len(s); i++ {
		canBreak := false
		for j := 0; j < i; j++ {
			if !canBreaks[j] {
				continue
			}
			if isInWordDict(s[j:i], wordDict) {
				canBreak = true
				break
			}
		}
		canBreaks[i] = canBreak
	}
	return canBreaks[len(s)]
}

func isInWordDict(s string, wordDict []string) bool {
	for _, word := range wordDict {
		if s == word {
			return true
		}
	}
	return false
}
