package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	leaves := getLeafPath(root)
	fmt.Printf("%v\n", leaves)
	sum := 0
	for _, leaf := range leaves {
		sum += getLeafNum(leaf)
	}
	return sum
}

func getLeafPath(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{[]int{root.Val}}
	}
	result := [][]int{}
	if root.Left != nil {
		leftLeafPath := getLeafPath(root.Left)
		for _, path := range leftLeafPath {
			result = append(result, append(path, root.Val))
		}
	}
	if root.Right != nil {
		rightLeafPath := getLeafPath(root.Right)
		for _, path := range rightLeafPath {
			result = append(result, append(path, root.Val))
		}
	}
	return result
}

func getLeafNum(nums []int) int {
	value := 0
	for i := 0; i < len(nums); i++ {
		value += nums[i] * int(math.Pow(float64(10), float64(i)))
	}
	return value
}
