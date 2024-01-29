package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, s := range strs {
		key := getKey(s)
		if groups, ok := groupMap[key]; ok {
			groupMap[key] = append(groups, s)
		} else {
			groupMap[key] = []string{s}
		}
	}
	fmt.Printf("%v", groupMap)
	result := make([][]string, 0)
	for _, groups := range groupMap {
		result = append(result, groups)
	}
	return result
}

func getKey(s string) string {
	slice := []rune{}
	for _, c := range s {
		slice = append(slice, c)
	}
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
	return string(slice)
}
