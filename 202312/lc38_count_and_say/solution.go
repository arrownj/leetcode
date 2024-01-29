package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	previousSay := countAndSay(n - 1)
	count := 1
	previousDigit := previousSay[0]
	say := ""
	for i := 1; i < len(previousSay); i++ {
		if previousSay[i] == previousDigit {
			count++
		} else {
			say += strconv.Itoa(count) + string(previousDigit)
			previousDigit = previousSay[i]
			count = 1
		}
	}
	say += strconv.Itoa(count) + string(previousDigit)
	return say
}
