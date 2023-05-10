package main

func RotateString(s, goal string) bool {
	if s == goal {
		return true
	}
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(goal); i++ {
		match := true
		for j := 0; j < len(s); j++ {
			if s[j] != goal[(i+j)%len(goal)] {
				match = false
				break
			}
		}
		if match {
			return match
		}
	}
	return false
}
