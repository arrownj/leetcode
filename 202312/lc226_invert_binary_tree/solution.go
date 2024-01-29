package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Invert(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	invertedLeft = Invert(root.Left)
	invertedRight = Invert(root.Right)
	root.Left = invertedLeft
	root.Right = invertedRight
	return root
}
