package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(MaxDepthRecursive(root.Left)), float64(MaxDepthRecursive(root.Right)))) + 1
}

func MaxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	depth := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		levelNodes := []*TreeNode{}
		for len(queue) > 0 {
			levelNodes = append(levelNodes, queue[0])
			queue = queue[1:]
		}
		depth += 1
		for _, node := range levelNodes {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
	return depth
}
