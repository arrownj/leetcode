package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			return [][]int{[]int{root.Val}}
		} else {
			return [][]int{}
		}
	}
	nodes := [][]int{}
	if root.Left != nil {
		leftNodes := PathSum(root.Left, targetSum-root.Val)
		if len(leftNodes) > 0 {
			for i := 0; i < len(leftNodes); i++ {
				newNodes := leftNodes[i]
				newNodes = append(newNodes[:1], newNodes[0:]...)
				newNodes[0] = root.Val
				nodes = append(nodes, newNodes)
			}
		}
	}
	if root.Right != nil {
		rightNodes := PathSum(root.Right, targetSum-root.Val)
		if len(rightNodes) > 0 {
			for i := 0; i < len(rightNodes); i++ {
				newNodes := rightNodes[i]
				newNodes = append(newNodes[:1], newNodes[0:]...)
				newNodes[0] = root.Val
				nodes = append(nodes, newNodes)
			}
		}
	}
	return nodes
}
