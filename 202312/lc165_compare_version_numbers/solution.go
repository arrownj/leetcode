package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1, version2 string) int {
	revision1 := strings.Split(version1, ".")
	revision2 := strings.Split(version2, ".")
	maxLength := max(len(revision1), len(revision2))
	for i := 0; i < maxLength; i++ {
		v1, v2 := 0, 0
		if i < len(revision1) {
			v1 = convert(revision1[i])
		}
		if i < len(revision2) {
			v2 = convert(revision2[i])
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convert(v string) int {
	value, _ := strconv.Atoi(v)
	return value
}
