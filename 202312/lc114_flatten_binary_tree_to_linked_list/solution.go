package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	previous := nil
	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if previous != nil {
			previous.Right = curr
			previous.Left = nil
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		previous = curr
	}
}
