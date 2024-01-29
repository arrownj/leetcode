package main

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i := 0
	for i < max(len(word1), len(word2)) {
		if i < len(word1) {
			result += word1[i : i+1]
		}
		if i < len(word2) {
			result += word2[i : i+1]
		}
		i++
	}
	return result
}
