package main

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == levelLength-1 {
				result = append(result, node.Val)
			}
		}
		queue = queue[levelLength:]
	}
	return result
}
