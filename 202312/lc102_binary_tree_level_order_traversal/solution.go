package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelLength := len(queue)
		levelValues := []int{}
		for i := 0; i < levelLength; i++ {
			levelValues = append(levelValues, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, levalValues)
		queue = queue[levelLength:]
	}
	return result
}
