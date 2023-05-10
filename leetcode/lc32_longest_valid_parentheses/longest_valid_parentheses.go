package main

func LongestValidParentheses(s string) int {
	longestLength := 0
	longestLengths := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == '(' {
			longestLengths[i] = 0
		} else { // s[i] == ")" && i >= 1
			if s[i-1] == '(' {
				if i >= 2 {
					longestLengths[i] = longestLengths[i-2] + 2
				} else {
					longestLengths[i] = 2
				}
			} else { // s[i-1] == ")"
				if i-longestLengths[i-1]-1 >= 0 {
					if s[i-longestLengths[i-1]-1] == '(' {
						if i-longestLengths[i-1]-2 >= 0 {
							longestLengths[i] = longestLengths[i-1] + longestLengths[i-longestLengths[i-1]-2] + 2
						} else {
							longestLengths[i] = longestLengths[i-1] + 2
						}
					} else {
						longestLengths[i] = 0
					}
				} else {
					longestLengths[i] = 0
				}
			}
		}

		if longestLength < longestLengths[i] {
			longestLength = longestLengths[i]
		}

	}
	return longestLength
}
