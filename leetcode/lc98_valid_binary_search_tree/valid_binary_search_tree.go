package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBinarySearchTree(root, math.MinInt, math.MaxInt)
}

func isValidBinarySearchTree(root *TreeNode, min, max int) bool {
	if root.Val <= min || root.Val >= max {
		return false
	}
	ret := true
	if root.Left != nil {
		ret = ret && isValidBinarySearchTree(root.Left, min, root.Val)
	}
	if root.Right != nil {
		ret = ret && isValidBinarySearchTree(root.Right, root.Val, max)

	}
	return ret
}
