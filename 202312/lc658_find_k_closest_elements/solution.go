package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	index := indexOf(arr, x, 0, len(arr)-1)
	fmt.Printf("index: %d\n", index)
	i, j := index, index+1
	for j-i-1 < k {
		if i < 0 {
			j++
		} else if j >= len(arr) {
			i--
		} else {
			if x-arr[i] <= arr[j]-x {
				i--
			} else {
				j++
			}
		}
	}
	return arr[i+1 : j]

}

func indexOf(arr []int, x int, left, right int) int {
	fmt.Printf("left: %d, right: %d\n", left, right)
	if left == right {
		if arr[left] >= x {
			return left - 1
		}
		return left
	}
	mid := left + (right-left)/2
	fmt.Printf("mid: %d\n", mid)
	if arr[mid] >= x {
		return indexOf(arr, x, left, mid)
	}
	return indexOf(arr, x, mid+1, right)

}
