package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}
	rootValue := preorder[0]
	rootIndex := 0
	for rootIndex < len(inorder) {
		if rootValue == inorder[rootIndex] {
			break
		}
	}
	leftInorder := inorder[0:rootIndex]
	rightInorder := inorder[rootIndex+1:]
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]
	return &TreeNode{
		Val:   rootValue,
		Left:  buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}

}
