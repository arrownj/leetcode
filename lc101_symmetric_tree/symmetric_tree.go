package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return IsSymmetric(root.Left, root.Right)

}

func IsSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return IsSymmetric(left.Left, right.Right) && IsSymmetric(left.Right, right.Left)
}
