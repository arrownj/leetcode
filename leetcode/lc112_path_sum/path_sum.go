package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}

	if root.Left != nil && HasPathSum(root.Left, sum-root.Val) {
		return true
	}

	if root.Right != nil && HasPathSum(root.Right, sum-root.Val) {
		return true
	}

	return false
}
