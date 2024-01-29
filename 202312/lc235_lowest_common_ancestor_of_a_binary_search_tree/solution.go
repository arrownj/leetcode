package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val <= p.Val && root.Val <= q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
