package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTreeRecursive(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	return p.Val == q.Val && isSameTreeRecursive(p.Left, q.Left) && isSameTreeRecursive(p.Right, q.Right)
}

func isSameTreeIterative(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	pQueue := []*TreeNode{p}
	qQueue := []*TreeNode{q}
	for len(pQueue) > 0 && len(qQueue) > 0 {
		pFirst := pQueue[0]
		qFirst := qQueue[0]
		pQueue = pQueue[1:]
		qQueue = qQueue[1:]
		if pFirst.Val != qFirst.Val {
			return false
		}
		if (pFirst.Left == nil && qFirst.Left != nil) || (pFirst.Left != nil && qFirst.Left == nil) {
			return false
		}
		if (pFirst.Right == nil && qFirst.Right != nil) || (pFirst.Right != nil && qFirst.Right == nil) {
			return false
		}
		if pFirst.Left != nil && qFirst.Left != nil {
			pQueue = append(pQueue, pFirst.Left)
			qQueue = append(qQueue, qFirst.Left)
		}
		if pFirst.Right != nil && qFirst.Right != nil {
			pQueue = append(pQueue, pFirst.Right)
			qQueue = append(qQueue, qFirst.Right)
		}
	}
	if len(pQueue) > 0 || len(qQueue) > 0 {
		return false
	}
	return true
}
