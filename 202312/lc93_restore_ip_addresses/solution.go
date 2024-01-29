package main

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	addresses := []string{}
	for i := 2; i <= len(s)-2; i++ {
		leftParts := getParts(s[:i])
		rightParts := getParts(s[i:])
		for _, l := range leftParts {
			for _, r := range rightParts {
				addresses = append(addresses, l+"."+r)
			}
		}
	}
	return addresses
}

func getParts(s string) []string {
	parts := []string{}
	for i := 1; i <= len(s)-1; i++ {
		left := s[:i]
		right := s[i:]
		if (left[0] == '0' && len(left) > 1) || (right[0] == '0' && len(right) > 1) {
			continue
		}
		leftInt, _ := strconv.Atoi(left)
		rightInt, _ := strconv.Atoi(right)
		if leftInt >= 0 && leftInt <= 255 && rightInt >= 0 && rightInt <= 255 {
			parts = append(parts, left+"."+right)
		}
	}
	return parts
}
