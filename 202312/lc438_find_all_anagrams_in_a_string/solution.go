package main

func findAnagrams(s, p string) []int {
	charMap := make(map[rune]bool)
	for _, r := range p {
		charMap[r] = true
	}
	result := []int{}
	i := 0
	for i < len(s) {
		matched := true
		for j := 0; j < len(p); j++ {
			if _, ok := charMap[s[i+j]]; !ok {
				i += j
				matched = false
				break
			} else {
			}
		}
		i += 1
	}

}
