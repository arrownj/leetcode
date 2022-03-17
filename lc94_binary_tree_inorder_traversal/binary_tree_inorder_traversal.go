package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	if root.Left != nil {
		ret = append(ret, InorderTraversal(root.Left)...)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		ret = append(ret, InorderTraversal(root.Right)...)
	}
	return ret
}
