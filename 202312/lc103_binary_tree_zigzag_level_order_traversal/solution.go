package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	result := [][]int{}
	leftToRight := true
	for len(queue) > 0 {
		levelLength := len(queue)
		levelValues := []int{}
		if leftToRight {
			for i := 0; i < levelLength; i++ {
				levelValues = append(levelValues, queue[i].Val)
			}
		} else {
			for i := levelLength - 1; i >= 0; i-- {
				levelValues = append(levelValues, queue[i].Val)
			}
		}
		for i := 0; i < levelLength; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, levelValues)
		leftToRight = !leftToRight
		queue = queue[levelLength:]
	}
	return result

}
