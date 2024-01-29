package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		path := []int{}
		root := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for root != nil {
			path = append(path, root.Val)
			root = root.Left
			queue = append(queue, root)
		}
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
	}
}
