package main

func isAnagram(s string, t string) bool {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if count, ok := sMap[c]; ok {
			sMap[c] = count + 1
		} else {
			sMap[c] = 1
		}
	}

	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		c := t[i]
		if count, ok := tMap[c]; ok {
			tMap[c] = count + 1
		} else {
			tMap[c] = 1
		}
	}

	for c, tCount := range tMap {
		if sCount, ok := sMap[c]; ok {
			if sCount != tCount {
				return false
			}
		} else {
			return false
		}
	}

	for c, sCount := range sMap {
		if tCount, ok := tMap[c]; ok {
			if sCount != tCount {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
