package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 1 {
		return true
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	diff := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}
