package main

func findTheDifference(s string, t string) byte {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if count, ok := sMap[c]; ok {
			sMap[c] = count + 1
		} else {
			sMap[c] = count + 1
		}
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		if count, ok := tMap[c]; ok {
			tMap[c] = count + 1
		} else {
			tMap[c] = count + 1
		}
	}

	result := t[0]
	for c, tCount := range tMap {
		if sCount, ok := sMap[c]; ok {
			if sCount < tCount {
				result = c
				break
			}
		} else {
			result = c
			break
		}
	}
	return result
}
