package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convert(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  convert(nums[:mid]),
		Right: convert(nums[mid+1:]),
	}
}
