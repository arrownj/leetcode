package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	trees := []*TreeNode{}
	for i := 1; i <= n; i++ {
		leftTrees := buildTree(1, i-1)
		rightTrees := buildTree(i+1, n)
		for l := 0; l < len(leftTrees); l++ {
			for r := 0; r < len(rightTrees); r++ {
				tree := &TreeNode{
					Val:   i,
					Left:  leftTrees[l],
					Right: rightTrees[r],
				}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}

func buildTree(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	if left == right {
		return []*TreeNode{&TreeNode{
			Val:   left,
			Left:  nil,
			Right: nil,
		}}
	}
	trees := []*TreeNode{}
	for i := left; i <= right; i++ {
		leftTrees := buildTree(left, i-1)
		rightTrees := buildTree(i+1, right)
		for l := 0; l < len(leftTrees); l++ {
			for r := 0; r < len(rightTrees); r++ {
				tree := &TreeNode{
					Val:   i,
					Left:  leftTrees[l],
					Right: rightTrees[r],
				}
				trees = append(trees, tree)
			}
		}
	}
	return trees
}
