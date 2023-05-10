package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return SameTree(t1.Left, t2.Left) && SameTree(t1.Right, t2.Right)
}
