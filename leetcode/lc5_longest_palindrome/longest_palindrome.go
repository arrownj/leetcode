package longest_palindrome

func GetLongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	var start, end = 0, 0
	for i := 0; i < len(s); i++ {
		start1, end1 := getLongestPalindromeWithCenter(s, i, i)
		if (end1 - start1 + 1) > (end - start + 1) {
			start = start1
			end = end1
		}
		if i < len(s)-1 {
			start2, end2 := getLongestPalindromeWithCenter(s, i, i+1)
			if (end2 - start2 + 1) > (end - start + 1) {
				start = start2
				end = end2
			}
		}
	}
	return s[start : end+1]
}

func getLongestPalindromeWithCenter(s string, start, end int) (int, int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start -= 1
		end += 1
	}
	return start + 1, end - 1
}
